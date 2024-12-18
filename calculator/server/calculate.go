package main

import (
	"context"

	pb "github.com/go-primer/calculator/proto"
)

func (s *Server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	num1 := req.Num1
	num2 := req.Num2

	return &pb.AddResponse{
		Result: num1 + num2,
	}, nil

}
