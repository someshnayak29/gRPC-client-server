package main

import (
	"fmt"
	"log"
	"net"

	"github.com/someshnayak29/grpc-server-client/chat"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Go gRPC Client Server Project!!!")
	// Creates a TCP listener on port:9000
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to Listen: %v", err)
	}

	// new gRPC server instance, we have to define it in chat.go
	s := chat.Server{}

	// creates a grpc server, at this pt. no services are registered to it
	grpcServer := grpc.NewServer()

	// Register the service implementation with grpc server
	chat.RegisterChatServiceServer(grpcServer, &s)

	// Start gRPC server which begins accepting & handling incoming req on this listener lis
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
