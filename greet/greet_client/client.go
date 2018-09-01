package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/octoberstorm/go-grpc-course/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello from Client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	// fmt.Printf("Created client: %f", c)

	// doUnary(c)
	// doServerStreaming(c)
	doClientStreaming(c)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	log.Println("Streaming...")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Joe",
			LastName:  "Dalton",
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error on GreetManyTimes: %v", err)
	}

	for {
		msg, err := resStream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error receiving streaming response: %v", err)
		}

		log.Printf("Response from GreetManyTimes: %v", msg.GetResult())
	}
}

func doUnary(c greetpb.GreetServiceClient) {
	log.Println("Start doing Unary request")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Joe",
			LastName:  "Dalton",
		},
	}

	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("Error requesting: %v", err)
	}

	log.Printf("Result from server: %v", res.Result)
}

func doClientStreaming(c greetpb.GreetServiceClient) {
	log.Println("Doing client streaming")

	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Joe",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Marry",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "John",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Mike",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Peter",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error creating client stream: %v", err)
	}

	for _, req := range requests {
		log.Printf("Sending request: %v\n", req)
		err := stream.Send(req)

		if err != nil {
			log.Fatalf("Error sending stream: %v", err)
		}

		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error receving streaming: %v", err)
	}

	log.Printf("Received ClientStream response: %v\n", res)
}
