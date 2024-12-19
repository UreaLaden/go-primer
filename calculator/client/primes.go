package main

import (
	"context"
	"io"
	"log"
	"strconv"
	"strings"

	pb "github.com/go-primer/calculator/proto"
)

func getPrimes(c pb.CalculatorServiceClient){
	log.Println("getPrimes was invoked")

	req := &pb.PrimesRequest{
		BaseNumber:120,
	}

	stream,err := c.Primes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling primes: %v\n",err)
	}
		
	var result []string
	for{
		msg, err := stream.Recv()
		if err == io.EOF{
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n",err)
		}

		convertedResult := strconv.Itoa(int(msg.Result))
		result = append(result,convertedResult)
	}
	resultString := "( " + strings.Join(result,", ") + " )"
	log.Printf("Primes: %s",resultString)
}