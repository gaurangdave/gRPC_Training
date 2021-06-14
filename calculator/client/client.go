package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"
	"training/calculator/calcpb"

	"google.golang.org/grpc"
)

func doMaxCall(c calcpb.CalculatorServiceClient) {
	fmt.Println("starting max number call")

	stream, err := c.Maximum(context.Background())

	if err != nil {
		log.Fatalf("Error creating the stream %v", err)
	}

	waitc := make(chan struct{})

	// sending bunch of numbers
	go func() {
		mynums := []int64{21, 21, 43, 23, 546, 8, 1, 0, 876}

		for _, num := range mynums {
			fmt.Printf("sending number %d\n", num)

			stream.Send(&calcpb.MaximumRequest{
				Num: num,
			})
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			resp, err := stream.Recv()

			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiveing %v", err)
				break
			}

			fmt.Printf("Max is %d\n", resp.Max)
		}
		close(waitc)
	}()

	<-waitc
}

func main() {
	fmt.Println("I am a calculator client...")

	clientConnection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect %v", err)
	}

	defer clientConnection.Close()

	// calling the unary RPC call
	serviceClient := calcpb.NewCalculatorServiceClient(clientConnection)

	req := &calcpb.CalculatorRequest{
		Num1: 12,
		Num2: 13,
	}

	res, err := serviceClient.Sum(context.Background(), req)

	if err != nil {
		log.Fatalf("error %v", err)
	}

	fmt.Printf("We got response from server : %v\n", res.Result)

	primeNumReq := &calcpb.PrimeNumDecompRequest{
		Num: 120,
	}

	resStr, primeErr := serviceClient.PrimeNumDecomposition(context.Background(), primeNumReq)

	if primeErr != nil {
		log.Fatalf("error calling the prime number decomposition sever%v", primeErr)
	}

	for {
		resp, err := resStr.Recv()

		if err == io.EOF {
			// end of stream
			fmt.Println("Thats all focks!")
			break
		}

		if err != nil {
			log.Fatalf("Error reading the stream  : %v", err)
		}

		fmt.Printf("%d\n", resp.Num)
	}

	reqStream, reqErr := serviceClient.Average(context.Background())

	if reqErr != nil {
		log.Fatalf("error calling the average number RPC : %v", reqErr)
	}

	for i := 1; i < 11; i++ {

		fmt.Printf("Sending %d\n", i)
		reqStream.Send(&calcpb.AverageRequest{
			Num: int64(i),
		})
	}

	averageResponse, respErr := reqStream.CloseAndRecv()

	if respErr != nil {
		fmt.Printf("Error getting average response %v", respErr)
		return
	}

	fmt.Printf("\nFinally got average response %v\n", averageResponse)

	doMaxCall(serviceClient)
}
