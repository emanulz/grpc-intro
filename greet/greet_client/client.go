package main

import (
	"context"
	"log"
	"time"

	pb "grpc/greet/greetpb"

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
	c := pb.NewGreetServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	greeting := &pb.Greeting{FirstName: "Emanuel", LastName: "Zuniga"}
	r, err := c.Greet(ctx, &pb.GreetRequest{Greeting: greeting})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("RESPONSE FROM SERVER: %s", r.GetResult())
}
