package main

import (
	"context"
	"log"

	pb "github.com/go-primer/calculator/proto"
)

func Add(c pb.CalculatorServiceClient) {
	log.Println("Add was invoked")

	req := &pb.AddRequest{
		Num1: 3,
		Num2: 10,
	}

	res, err := c.Add(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not calculate: %v\n", err)
	}

	log.Printf("%d + %d = %d\n", req.Num1, req.Num2, res.Result)
}
