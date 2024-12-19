package main

import (
	"context"
	"log"
	"time"

	pb "github.com/go-primer/calculator/proto"
)

func doAverages(c pb.CalculatorServiceClient) {
	log.Println("doAverages function was invoked")

	reqs := []*pb.AverageRequest{
		{Value: 1},
		{Value: 2},
		{Value: 3},
		{Value: 4},
	}

	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Average: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving repsonse from Average: %v\n", err)
	}

	log.Printf("Average: (%f)", res.Average)
}
