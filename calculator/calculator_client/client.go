package main

import (
	"context"
	"io"
	"log"

	"github.com/octoberstorm/go-grpc-course/calculator/calculatorpb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Error dialing: %v", err)
	}

	defer conn.Close()

	cc := calculatorpb.NewCalculatorServiceClient(conn)

	// duUnary(cc)
	// doServerStreaming(cc)
	doClientStreaming(cc)
}

func doServerStreaming(cc calculatorpb.CalculatorServiceClient) {
	log.Println("Streaming...")

	req := &calculatorpb.PrimeNumberRequest{
		Number: 120,
	}

	resStream, err := cc.PrimeStream(context.Background(), req)

	if err != nil {
		log.Fatalf("Error streaming: %v", err)
	}

	for {
		msg, err := resStream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error recv: %v", err)
		}

		log.Printf("%v", msg.GetPrime())
	}
}

func duUnary(cc calculatorpb.CalculatorServiceClient) {
	log.Println("Start unary request")

	req := &calculatorpb.SumRequest{
		Sum: &calculatorpb.Sum{
			FirstNumber:  1,
			SecondNumber: 2,
		},
	}

	res, err := cc.Sum(context.Background(), req)

	if err != nil {
		log.Fatalf("Error calculating: %v", err)
	}

	log.Printf("Result: %v", res.Result)
}

func doClientStreaming(cc calculatorpb.CalculatorServiceClient) {
	log.Println("Start client streaming requests...")

	requests := []*calculatorpb.ComputeAverageRequest{
		&calculatorpb.ComputeAverageRequest{
			Number: 1,
		},
		&calculatorpb.ComputeAverageRequest{
			Number: 2,
		},
		&calculatorpb.ComputeAverageRequest{
			Number: 3,
		},
		&calculatorpb.ComputeAverageRequest{
			Number: 4,
		},
	}

	stream, err := cc.ComputeAverage(context.Background())

	for _, req := range requests {
		log.Printf("Sending request: %v\n", req)

		stream.Send(req)
		// time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	log.Printf("Average: %v", res.GetAverage())
}
