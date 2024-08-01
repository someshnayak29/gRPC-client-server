# gRPC Chat Service

This project demonstrates a simple gRPC chat service using TCP communication between a gRPC client and server. The server listens for requests on port 9000 and responds to messages sent by the client.

## Project Structure


## Prerequisites

- Go 1.18 or higher
- Protocol Buffers compiler (`protoc`)
- `protoc-gen-go` and `protoc-gen-go-grpc` plugins for Go

## Setup

### 1. Install Dependencies

Make sure you have the required Go modules installed. Run the following command in the project root directory:

```sh
go mod tidy

# Install the Protocol Buffers compiler (protoc)
# For macOS
brew install protobuf

# For Ubuntu
sudo apt-get install -y protobuf-compiler

# Install the Go plugins for Protocol Buffers
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# gRPC Chat Service

This project demonstrates a simple gRPC chat service using TCP communication between a gRPC client and server. The server listens for requests on port 9000 and responds to messages sent by the client.

## Project Structure

grpc-chat/
├── chat/
│ ├── chat.proto # The Protocol Buffers definition
│ ├── chat.pb.go # Generated Go code for Protocol Buffers
│ └── chat_grpc.pb.go # Generated Go code for gRPC service
├── client/
│ └── client.go # The gRPC client implementation
├── server.go # The gRPC server implementation
└── go.mod # Go module file

### cd into folder which has proto file
cd chat
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative chat.proto

cd ..
go run main.go

### Running the Client
To test the server, navigate to the client/ directory and run:

cd client
go run main.go

###Troubleshooting

Best practice to have chat.proto in same package as chat.go and this package should be mentioned in options go_package = "path/chat" in chat.go

Server struct in chat.go should use UnimplementedChatServiceServer, it's a way to ensure that your server type implements the necessary interface for gRPC service handling


