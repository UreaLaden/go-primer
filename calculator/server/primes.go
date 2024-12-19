package main

import (
	"log"

	pb "github.com/go-primer/calculator/proto"
	"google.golang.org/grpc"
)

func (s *Server) Primes(in *pb.PrimesRequest, stream grpc.ServerStreamingServer[pb.PrimesResponse]) error {
	log.Printf("Prims function was invoked with: %v\n", in)

	k := uint32(2)
	n := in.BaseNumber

	for n > 1 {
		if n%k == 0 {
			stream.Send(&pb.PrimesResponse{
				Result: k,
			})
			n = n / k
		}else{
			k++
		}
	}
	return nil
}
