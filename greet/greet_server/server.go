package greet

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"training/greet/greetpb"

	"google.golang.org/grpc"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func (server) Greet(ctx context.Context, req *greetpb.GreetReqeust) (*greetpb.GreetResponse, error) {
	fname := req.GetGreeting().GetFirstName()
	result := "Hello " + fname
	res := greetpb.GreetResponse{
		Result: result,
	}

	return &res, nil
}

func (server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fname := req.GetGreeting().GetFirstName()
	result := "Hello " + fname
	for i := 0; i < 10; i++ {
		stream.Send(&greetpb.GreetManyTimesResponse{
			Result: result,
		})
		time.Sleep(1000 * time.Millisecond)
	}

	return nil

}

// function to handle client streams
func (server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	result := ""
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			// we received end of stream from client
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading the client stream %v", err)
		}

		name := req.Name
		result += "Hello " + name + "!"
	}

	return nil
}

func (server) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		name := req.Name

		stream.Send(&greetpb.GreetEveryoneResponse{
			Message: "Hello! " + name,
		})

	}
}

// func (server) mustEmbedUnimplementedGreetServiceServer() {}

func Run() {
	fmt.Println("starting greet server...")

	// create listener
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()
	myser := &server{}

	greetpb.RegisterGreetServiceServer(s, myser)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
