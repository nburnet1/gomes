package main


import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/golang-jwt/jwt/v5"
	pb "github.com/nburnet1/gomes/proto"
	"google.golang.org/grpc"

	"github.com/nburnet1/gomes/internal/auth"
)

var users = map[string]string{
	"admin": "password123",
	"user":  "password123",
}

var roles = map[string]string{
	"admin": "admin",
	"user":  "user",
}


var jwtSecret = []byte("supersecretkey")

type AuthServiceServer struct {
	pb.UnimplementedAuthServiceServer
}

func (s *AuthServiceServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	// Validate credentials
	if pass, exists := users[req.Username]; !exists || pass != req.Password {
		return nil, fmt.Errorf("invalid credentials")
	}

	// Create JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": req.Username,
		"role":     roles[req.Username],
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, _ := token.SignedString(jwtSecret)
	return &pb.LoginResponse{Token: tokenString}, nil
}

func (s *AuthServiceServer) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(req.Token, &claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return &pb.ValidateTokenResponse{Valid: false}, nil
	}

	return &pb.ValidateTokenResponse{
		Valid:    true,
		Username: claims["username"].(string),
		Role:     claims["role"].(string),
	}, nil
}



func main() {

	auth.InitDB()

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, &AuthServiceServer{})
	log.Println("Auth Service running on port 50051")
	grpcServer.Serve(listener)
}