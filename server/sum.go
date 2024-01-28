package main

import (
	"context"
	"log"

	pb "grpc-calculator/proto"
)

func (*Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum was invoked with %v\n", in)
	return &pb.SumResponse{Result: in.FirstNumber + in.SecondNumber}, nil
}
