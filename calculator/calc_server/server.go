package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"
	"training/calculator/calcpb"

	"google.golang.org/grpc"
)

type server struct {
	calcpb.UnimplementedCalculatorServiceServer
}

func (server) Sum(ctx context.Context, req *calcpb.CalculatorRequest) (*calcpb.CalculatorResponse, error) {
	num1 := req.Num1
	num2 := req.Num2

	return &calcpb.CalculatorResponse{
		Result: num1 + num2,
	}, nil
}

func (server) PrimeNumDecomposition(req *calcpb.PrimeNumDecompRequest, stream calcpb.CalculatorService_PrimeNumDecompositionServer) error {
	num := req.Num
	k := int64(2)
	for num > 1 {
		if num%k == 0 {
			stream.Send(&calcpb.PrimeNumDecompResponse{Num: k})
			time.Sleep(1000 * time.Millisecond)
			num = num / k
		} else {
			k = k + 1
		}
	}

	return nil
}

func (server) Average(stream calcpb.CalculatorService_AverageServer) error {
	/**
	    1. listen to stream
	    2. calculate sum
	    3. calculate number of input
	**/
	sum := int64(0)
	total := int64(1)

	for {
		reqStr, reqErr := stream.Recv()

		if reqErr == io.EOF {
			// end of stream - send resonse
			return stream.SendAndClose(&calcpb.AverageResponse{
				Result: sum / total,
			})
		}

		if reqErr != nil {
			log.Fatalf("Error while reading the client stream %v", reqErr)
		}

		num := reqStr.Num
		sum = sum + num
		total = total + 1
	}

	return nil
}

func (server) Maximum(stream calcpb.CalculatorService_MaximumServer) error {
	var previousMax int64 = 0
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			// send max of all numbers at the end of the stream.
			return nil
		}

		if err != nil {
			log.Fatalf("Error receivng the stream %v", err)
			return err
		}

		nextNum := req.GetNum()
		if nextNum > previousMax {
			previousMax = nextNum
			stream.Send(
				&calcpb.MaximumResponse{
					Max: previousMax,
				},
			)
		}

	}

	return nil
}

func Run() {
	fmt.Println("starting calculator server...")

	// create listener
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()
	myser := &server{}

	calcpb.RegisterCalculatorServiceServer(s, myser)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}

}
