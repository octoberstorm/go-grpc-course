package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/octoberstorm/go-grpc-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	log.Printf("Received request: %v", req)

	sum := req.GetSum()
	firstNumber := sum.GetFirstNumber()
	secondNumber := sum.GetSecondNumber()

	res := &calculatorpb.SumResponse{
		Result: firstNumber + secondNumber,
	}

	return res, nil
}

func (*server) PrimeStream(req *calculatorpb.PrimeNumberRequest, stream calculatorpb.CalculatorService_PrimeStreamServer) error {
	number := req.GetNumber()

	log.Printf("Streaming prime numbers for %v", number)
	var k int32 = 2

	for {
		if number%k == 0 {
			log.Printf("Factor: %v", k)
			res := &calculatorpb.PrimeNumberResponse{
				Prime: k,
			}
			stream.Send(res)
			time.Sleep(1000 * time.Millisecond)
			number = number / k
		} else {
			k = k + 1
		}

		if number <= 1 {
			break
		}
	}

	return nil
}

// for client streaming
func (*server) ComputeAverage(stream calculatorpb.CalculatorService_ComputeAverageServer) error {
	log.Println("Start client streaming...")

	var average float32
	var count, sum int32

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&calculatorpb.ComputeAverageResponse{
				Average: average,
			})
		}

		if err != nil {
			log.Fatalf("Error receving client stream: %v", err)
		}

		count++
		sum += req.GetNumber()
		average = float32(float32(sum) / float32(count))
	}
}

func main() {
	fmt.Println("Starting server on port 50051")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	s := grpc.NewServer()

	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error serving: %v", err)
	}
}
