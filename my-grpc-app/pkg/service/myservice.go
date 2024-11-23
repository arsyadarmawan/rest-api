package service

import (
	"context"
	pb "my-grpc-app/api/proto"
)

// MyServiceServer implements the gRPC MyServiceServer interface.
type MyServiceServer struct {
	pb.UnimplementedMyServiceServer
}

// SayHello is the method that will be called by the gRPC client.
func (s *MyServiceServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello, " + req.Name}, nil
}
