package main

import pb "grpc-calculator/proto"

type Server struct {
	pb.CalculatorServiceServer
}
