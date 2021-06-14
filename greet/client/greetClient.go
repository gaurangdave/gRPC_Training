package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"training/greet/greetpb"

	"google.golang.org/grpc"
)

func doBiDiStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("starting BiDi Streaming RPC")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while creating the stream : %v", err)
		return
	}

	waitc := make(chan struct{})
	
	// sending bunch of messages
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("sending message", fmt.Sprintf("Jonny_%d", i))
			stream.Send(&greetpb.GreetEveryoneRequest{
				Name: fmt.Sprintf("Jonny_%d", i),
			})
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	// func to receive messages
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

			fmt.Printf("%s\n", resp.Message)
		}
		close(waitc)
	}()

	<-waitc
}

func main() {
	fmt.Println("hello world, I am the client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	req := &greetpb.GreetReqeust{
		Greeting: &greetpb.Greeting{
			FirstName: "aa",
			LastName:  "bb",
		},
	}

	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("error %v", err)
	}

	fmt.Printf("We got response from server : %v\n", res.Result)

	manyTimesReq := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Tony",
			LastName:  "Stark",
		},
	}

	// what is context.Background?
	mtc, mtcError := c.GreetManyTimes(context.Background(), manyTimesReq)

	if mtcError != nil {
		log.Fatalf("error %v", mtcError)
	}

	for {
		resp, err := mtc.Recv()

		if err == io.EOF {
			// we've reached end of stream
			break
		}

		if err != nil {
			// some other error
			log.Fatalf(" error%v", err)
		}

		fmt.Printf("We got response from server : %v\n", resp.Result)
	}

	stream, longGreetErr := c.LongGreet(context.Background())

	if longGreetErr != nil {
		log.Fatalf("Error requesting long greet %v", longGreetErr)
	}

	for i := 0; i < 10; i++ {
		stream.Send(&greetpb.LongGreetRequest{
			Name: fmt.Sprintf("foo_%d", i),
		})
	}

	resp, streamError := stream.CloseAndRecv()

	if streamError != nil {
		log.Fatalf("Error requesting long greet %v", streamError)
	}

	fmt.Printf("Got Response : %s", resp)

	doBiDiStreaming(c)
}
