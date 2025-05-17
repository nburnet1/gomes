package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/nburnet1/gomes/internal/config"
	"github.com/nburnet1/gomes/pkg/namespace"
	pb "github.com/nburnet1/gomes/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// NamespaceServer implements NamespaceService
type NamespaceServer struct {
	pb.UnimplementedNamespaceServiceServer
	namespaceEngine *namespace.NamespaceEngine
	subscribers     map[string][]chan *pb.NodeUpdate
}

// CreateNode RPC
func (s *NamespaceServer) CreateNode(ctx context.Context, in *pb.CreateNodeRequest) (*pb.Node, error) {
	node, err := s.namespaceEngine.CreateNode(in.Topic, in.Value, nil)
	if err != nil {
		return nil, err
	}
	return &pb.Node{
		Topic:       node.GetTopic(),
		Value:       node.GetValue(),
		ParentTopic: node.GetParentTopic(),
		Name:        node.GetName(),
		Timestamp:   node.GetTimeStamp().String(),
	}, nil
}

// ReadNode RPC
func (s *NamespaceServer) ReadNode(ctx context.Context, in *pb.ReadNodeRequest) (*pb.Node, error) {
	node, err := s.namespaceEngine.ReadNode(in.Topic)
	if err != nil {
		return nil, err
	}

	return &pb.Node{
		Topic:       node.GetTopic(),
		Value:       node.GetValue(),
		ParentTopic: node.GetParentTopic(),
		Name:        node.GetName(),
		Timestamp:   node.GetTimeStamp().String(),
	}, nil
}

// GetChildren RPC
func (s *NamespaceServer) GetChildren(ctx context.Context, in *pb.ReadNodeRequest) (*pb.BrowseNodesResponse, error) {
	node, err := s.namespaceEngine.ReadNode(in.Topic)
	if err != nil {
		return nil, err
	}

	children := node.GetChildren()
	pbNodes := []*pb.Node{}

	for _, child := range children {
		pbNodes = append(pbNodes, &pb.Node{
			Topic:       child.GetTopic(),
			Value:       child.GetValue(),
			ParentTopic: child.GetParentTopic(),
			Name:        child.GetName(),
			Timestamp:   child.GetTimeStamp().String(),
		})
	}

	return &pb.BrowseNodesResponse{Nodes: pbNodes}, nil
}

// UpdateNode RPC
func (s *NamespaceServer) UpdateNode(ctx context.Context, in *pb.UpdateNodeRequest) (*pb.Node, error) {
	node, err := s.namespaceEngine.UpdateNode(in.Topic, in.Value)
	if err != nil {
		return nil, err
	}


	return &pb.Node{
		Topic:       node.GetTopic(),
		Value:       node.GetValue(),
		ParentTopic: node.GetParentTopic(),
		Name:        node.GetName(),
		Timestamp:   node.GetTimeStamp().String(),
	}, nil
}


// DeleteNode RPC
func (s *NamespaceServer) DeleteNode(ctx context.Context, in *pb.DeleteNodeRequest) (*pb.DeleteNodeResponse, error) {
	err := s.namespaceEngine.DeleteNode(in.Topic)
	if err != nil {
		return &pb.DeleteNodeResponse{Success: false}, err
	}
	return &pb.DeleteNodeResponse{Success: true}, nil
}

// BrowseNodes RPC
func (s *NamespaceServer) BrowseNodes(ctx context.Context, in *pb.BrowseNodesRequest) (*pb.BrowseNodesResponse, error) {
	nodes := s.namespaceEngine.BrowseNodes()
	pbNodes := []*pb.Node{}

	for _, node := range nodes {
		pbNodes = append(pbNodes, &pb.Node{
			Topic:       node.GetTopic(),
			Value:       node.GetValue(),
			ParentTopic: node.GetParentTopic(),
			Name:        node.GetName(),
			Timestamp:   node.GetTimeStamp().String(),
		})
	}
	return &pb.BrowseNodesResponse{Nodes: pbNodes}, nil
}


// BrowseRootNodes RPC
func (s *NamespaceServer) BrowseRootNodes(ctx context.Context, in *pb.BrowseRootNodesRequest) (*pb.BrowseNodesResponse, error) {
	nodes := s.namespaceEngine.BrowseRootNodes()
	pbNodes := []*pb.Node{}
	
	for _, node := range nodes {
		pbNodes = append(pbNodes, &pb.Node{
			Topic:       node.GetTopic(),
			Value:       node.GetValue(),
			ParentTopic: node.GetParentTopic(),
			Name:        node.GetName(),
			Timestamp:   node.GetTimeStamp().String(),
		})
	}

	return &pb.BrowseNodesResponse{Nodes: pbNodes}, nil
}

// SubNode (streaming RPC)
func (s *NamespaceServer) SubNode(in *pb.SubNodeRequest, stream pb.NamespaceService_SubNodeServer) error {
	// Subscribe to NamespaceEngine
	updateChan, err := namespace.Engine.SubNode(in.Topic)
	if err != nil {
		return fmt.Errorf("failed to subscribe to node: %v", err)
	}

	// Stream updates from NamespaceEngine to gRPC clients
	for update := range updateChan {
		nodeUpdate := &pb.NodeUpdate{
			Topic: in.Topic,
			Value: update, // Convert the value to a string
		}

		if err := stream.Send(nodeUpdate); err != nil {
			log.Printf("Error sending update for %s: %v", in.Topic, err)
			return err
		}
	}

	return nil
}

// Get parent topic (if exists)

func main() {
	// Load configuration
	if err := config.MigrateFromRegistry(); err != nil {
		fmt.Println("Config migration error:", err)
	}


	// Start Namespace Engine
	if err := namespace.Engine.StartFromRegistry(); err != nil {
		fmt.Println("Namespace engine startup error:", err)
	}

	// Create test nodes
	namespace.Engine.CreateNode("Enterprise/someTag", map[string]string{"hello": "hello"}, nil)
	namespace.Engine.CreateNode("anotherTag", "test", nil)
	namespace.Engine.CreateNode("SomeNodeToTest", 45, nil)
	_, err := namespace.Engine.CreateNode("Enterprise/Site/Area/Machine1/AnotherNode", 4510, nil)
	namespace.Engine.CreateNode("Enterprise/Site/testTag", "test", nil)
	if err != nil {
		fmt.Println("Error creating node:", err)
	}


	fmt.Println("nodes:",namespace.Engine.BrowseNodes())
	// gRPC Server
	server := grpc.NewServer()
	namespaceServer := &NamespaceServer{
		namespaceEngine: namespace.Engine,
		subscribers:     make(map[string][]chan *pb.NodeUpdate),
	}
	pb.RegisterNamespaceServiceServer(server, namespaceServer)
	reflection.Register(server)
	listener, err := net.Listen("tcp", ":50055")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("Namespace gRPC service running on :50055")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	
}
