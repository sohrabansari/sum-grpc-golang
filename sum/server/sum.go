package main

import (
	"context"
	"log"

	pb "github.com/sohrabansari/sum-grpc/sum/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum was invoked with two numbers: %d & %d\n", in.Num1, in.Num2)
	return &pb.SumResponse{
		Result: in.Num1 + in.Num2,
	}, nil
}
