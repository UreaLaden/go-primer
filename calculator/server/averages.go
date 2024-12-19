package main

import (
	"io"
	"log"

	pb "github.com/go-primer/calculator/proto"
	"google.golang.org/grpc"
)

func (s *Server) Average(stream grpc.ClientStreamingServer[pb.AverageRequest, pb.AverageResponse]) error {
	log.Println("Average function was invoked")

	var values []int64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			average := getAverage(values)
			return stream.SendAndClose(&pb.AverageResponse{Average: average})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
			return err
		}

		log.Printf("Receiving: %v\n", req)

		values = append(values, req.Value)
	}
}

func getAverage(values []int64) float64 {
	if len(values) == 0 {
		return 0
	}

	var sum int64
	for _, val := range values {
		sum += val
	}

	var average float64
	average = float64(sum) / float64(len(values))
	return average
}
