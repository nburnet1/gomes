package grpcclient

import (
	"log"
	"sync"

	pb "github.com/nburnet1/gomes/proto"
	"google.golang.org/grpc"
)

var (
	AuthClient pb.AuthServiceClient
	AuthGrpcConn        *grpc.ClientConn
	authonce            sync.Once
)

// initGRPC initializes the gRPC connection only once
func InitAuthGRPC() {
	authonce.Do(func() {
		var err error
		AuthGrpcConn, err = grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to connect to gRPC server: %v", err)
		}
		AuthClient = pb.NewAuthServiceClient(AuthGrpcConn)
	})
}
