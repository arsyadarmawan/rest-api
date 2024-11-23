package main

import (
	"google.golang.org/grpc"
	"log"
	pb "my-grpc-app/api/proto"
	"my-grpc-app/pkg/service"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// Register the service
	pb.RegisterMyServiceServer(grpcServer, &service.MyServiceServer{})

	log.Println("gRPC server is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
