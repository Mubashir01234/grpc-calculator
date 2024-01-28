package main

import (
	"context"
	"fmt"
	"log"
	"math"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "grpc-calculator/proto"
)

func (*Server) Sqrt(ctx context.Context, req *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("Sqrt was invoked with number %d\n", req.Number)

	number := req.Number

	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %d", number),
		)
	}

	return &pb.SqrtResponse{
		Result: math.Sqrt(float64(number)),
	}, nil
}
