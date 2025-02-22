package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"sync"

	pb "github.com/nburnet1/gomes/proto"

	"github.com/fsnotify/fsnotify"
	"google.golang.org/grpc"
)

// ConfigServer implements the ConfigService gRPC API.
type ConfigServer struct {
	pb.UnimplementedConfigServiceServer
	configData map[string]map[string]string // Store configs in memory
	mu        sync.RWMutex
}

// LoadConfig reads the config file and updates in-memory storage.
func (s *ConfigServer) LoadConfig() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	file, err := os.Open("./cmd/config/config.json")
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	var configs map[string]map[string]string
	if err := json.Unmarshal(data, &configs); err != nil {
		return err
	}

	s.configData = configs
	log.Println("Configurations reloaded successfully.")
	return nil
}

// WatchConfigFile watches the config file for changes and reloads dynamically.
func (s *ConfigServer) WatchConfigFile() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	configFile := "./cmd/config/config.json"
	err = watcher.Add(configFile)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Watching config file for changes...")
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Write == fsnotify.Write {
				log.Println("Config file changed, reloading...")
				if err := s.LoadConfig(); err != nil {
					log.Println("Error reloading config:", err)
				}
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("Watcher error:", err)
		}
	}
}

// GetConfig retrieves the configuration for a service.
func (s *ConfigServer) GetConfig(ctx context.Context, req *pb.GetConfigRequest) (*pb.GetConfigResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	config, exists := s.configData[req.ServiceName]
	if !exists {
		return &pb.GetConfigResponse{Config: map[string]string{}}, fmt.Errorf("config not found for service: %s", req.ServiceName)
	}

	return &pb.GetConfigResponse{Config: config}, nil
}

func main() {
	server := grpc.NewServer()
	configServer := &ConfigServer{configData: make(map[string]map[string]string)}

	// Load initial config
	if err := configServer.LoadConfig(); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Start watching config file for changes
	go configServer.WatchConfigFile()

	pb.RegisterConfigServiceServer(server, configServer)

	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("Config service running on :50052")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}