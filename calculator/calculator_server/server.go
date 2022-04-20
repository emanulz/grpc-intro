package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grpc/calculator/calculatorpb"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedSumServiceServer
}

func (s *server) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	fmt.Println("Request for sum API recieved")
	num_1 := req.GetNum_1()
	num_2 := req.GetNum_2()
	return &pb.SumResponse{Result: num_1 + num_2}, nil
}

func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	// create a server and listen
	s := grpc.NewServer()
	pb.RegisterSumServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
