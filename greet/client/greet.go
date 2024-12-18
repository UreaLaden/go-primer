package main

import (
	"context"
	"log"

	pb "github.com/go-primer/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invooked")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Leaundrae",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}
