package main

import (
	"context"
	"log"

	"github.com/someshnayak29/grpc-server-client/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	var conn *grpc.ClientConn

	// establishes the connection to the new server
	conn, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials())) // insecure means we are not using proper encryption now
	if err != nil {
		log.Fatalf("Did not connect : %s", err)
	}
	// to ensure connection closes when function exists
	defer conn.Close()

	// This function creates a new client for the 'ChatService' service. Client is used to make RPC calls to the server
	c := chat.NewChatServiceClient(conn)

	// making gRPC call
	response, err := c.SayHello(context.Background(), &chat.Message{Body: "Hello from Client!!!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)
}
