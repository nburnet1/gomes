package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	pb "github.com/nburnet1/gomes/proto"
	"google.golang.org/grpc"
)

// MQTT settings
const (
	mqttBroker = "tcp://localhost:1883" // Change to your broker's address
	grpcServer = "localhost:50055"      // gRPC server address
)

// Connects to MQTT
func connectMQTT() mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(mqttBroker)
	opts.SetClientID("namespace_mqtt_client")
	opts.SetKeepAlive(60 * time.Second)
	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		log.Printf("MQTT Message Received: %s -> %s", msg.Topic(), string(msg.Payload()))
	})

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("MQTT Connection Failed: %v", token.Error())
	}
	log.Println("Connected to MQTT broker")
	return client
}


// Subscribe to gRPC stream for node updates
func subscribeToNodeUpdates(client mqtt.Client) {
	// Connect to gRPC server and keep it open
	conn, err := grpc.Dial(grpcServer, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	grpcClient := pb.NewNamespaceServiceClient(conn)

	// Fetch all nodes
	browseResp, err := grpcClient.BrowseNodes(context.Background(), &pb.BrowseNodesRequest{})
	if err != nil {
		log.Fatalf("Failed to browse nodes: %v", err)
	}

	// Subscribe to each node
	for _, node := range browseResp.Nodes {
		go func(topic string) {
			req := &pb.SubNodeRequest{Topic: topic}
			stream, err := grpcClient.SubNode(context.Background(), req)
			if err != nil {
				log.Printf("Failed to subscribe to node %s: %v", topic, err)
				return
			}

			// Receive updates and publish to MQTT
			for {
				update, err := stream.Recv()
				if err != nil {
					log.Printf("Stream closed for %s: %v", topic, err)
					return
				}

				// Convert update to JSON
				jsonValue, err := json.Marshal(update)
				if err != nil {
					log.Printf("Failed to encode JSON: %v", err)
					continue
				}

				// Publish update to MQTT
				token := client.Publish(update.Topic, 0, false, jsonValue)
				token.Wait()
				log.Printf("Published update to MQTT: %s -> %s", update.Topic, string(jsonValue))
			}
		}(node.Topic)
	}

	// Keep gRPC connection alive
	select {} // Blocks forever
}

func main() {
	// Start MQTT client
	mqttClient := connectMQTT()
	defer mqttClient.Disconnect(250)

	// Subscribe to gRPC updates and publish to MQTT
	subscribeToNodeUpdates(mqttClient)

	// Keep the program running
	select {}
}