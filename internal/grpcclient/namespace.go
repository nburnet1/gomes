package grpcclient

import (
	"log"
	"sync"

	pb "github.com/nburnet1/gomes/proto"
	"google.golang.org/grpc"
)

var (
	NamespaceClient pb.NamespaceServiceClient
	GrpcConn        *grpc.ClientConn
	once            sync.Once
)

// initGRPC initializes the gRPC connection only once
func InitGRPC() {
	once.Do(func() {
		var err error
		GrpcConn, err = grpc.Dial("localhost:50055", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to connect to gRPC server: %v", err)
		}
		NamespaceClient = pb.NewNamespaceServiceClient(GrpcConn)
	})
}
