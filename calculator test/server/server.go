package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "test.com/grpctest/proto"
)

const (
	port = ":50051"
)

type CalculatorServer struct {
	pb.UnimplementedCalculatorServer
}

func (s *CalculatorServer) Add(ctx context.Context, req *pb.AddRequest) (*pb.Response, error) {
	result := req.N1 + req.N2
	return &pb.Response{Response: result}, nil
}

func (s *CalculatorServer) Subtract(ctx context.Context, req *pb.SubtractRequest) (*pb.Response, error) {
	result := req.N1 - req.N2
	return &pb.Response{Response: result}, nil
}

func (s *CalculatorServer) Multiply(ctx context.Context, req *pb.MultiplyRequest) (*pb.Response, error) {
	result := req.N1 * req.N2
	return &pb.Response{Response: result}, nil
}

func (s *CalculatorServer) Divide(ctx context.Context, req *pb.DivideRequest) (*pb.Response, error) {
	result := req.N1 / req.N2
	return &pb.Response{Response: result}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &CalculatorServer{})
	log.Printf("server listening on %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}
