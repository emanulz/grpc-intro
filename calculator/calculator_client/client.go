package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "grpc/calculator/calculatorpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSumServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	sumMessage := &pb.SumRequest{Num_1: -25, Num_2: 14}
	fmt.Println("Sending request to server")
	r, err := c.Sum(ctx, sumMessage)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("RESPONSE FROM SERVER: %v", r.GetResult())
}
