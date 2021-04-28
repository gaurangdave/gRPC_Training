package greet

import (
	"fmt"
	"log"
	"net"

	"training/greet/greetpb"

	"google.golang.org/grpc"
)

// type server struct{}

// func (*server) mustEmbedUnimplementedGreetServiceServer() {
// 	fmt.Println("greeting service")
// }

func Run() {
	fmt.Println("Hello World")

	// create listener
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()
	myser := &greetpb.UnimplementedGreetServiceServer{}

	greetpb.RegisterGreetServiceServer(s, myser)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
