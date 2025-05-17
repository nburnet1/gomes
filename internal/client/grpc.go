package client

import (
	"log"
	"sync"

	pb "github.com/nburnet1/gomes/proto"
	"google.golang.org/grpc"
)

var (
	AuthGRPC      = &GRPCClient[pb.AuthServiceClient]{
		initClient: pb.NewAuthServiceClient,
	}
	NamespaceGRPC = &GRPCClient[pb.NamespaceServiceClient]{
		initClient: pb.NewNamespaceServiceClient,
	}

)


// GRPCClient is a generic struct to hold a gRPC client and connection
type GRPCClient[T any] struct {
	conn   *grpc.ClientConn
	Client T
	once   sync.Once
	initClient func(grpc.ClientConnInterface) T
}

// NewGRPCClient initializes a new GRPCClient instance
func NewGRPCClient[T any]() *GRPCClient[T] {
	return &GRPCClient[T]{}
}

// Connect establishes a gRPC connection (only once) and initializes the service client
func (c *GRPCClient[T]) Connect(address string) {
	c.once.Do(func() {
		var err error
		c.conn, err = grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to connect to gRPC server at %s: %v", address, err)
		}
		c.Client = c.initClient(c.conn) // Use the provided function to create the client
	})
}

// Close closes the gRPC connection
func (c *GRPCClient[T]) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}