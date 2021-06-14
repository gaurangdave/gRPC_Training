package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	donutspb "training/donuts/protos"
)

// Function to make AreYouOpen GPRC Method
func doAreYouOpen(client donutspb.DonutsServiceClient) {
	resp, err := client.AreYouOpen(context.Background(), &donutspb.AreYouOpenRequest{
		DayOfTheWeek: "Tuesday",
	})

	if err != nil {
		log.Fatalf("got error from donuts server %v", err)
		return
	}

	fmt.Printf("\ngot response from donuts server : %v\n", resp.Reponse)
}

func main() {
	fmt.Println("starting donuts client...")

	clientConnection, clientConnectionError := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if clientConnectionError != nil {
		log.Fatalf("error connecting to grpc server %v", clientConnectionError)
	}

	defer clientConnection.Close()

	serviceClient := donutspb.NewDonutsServiceClient(clientConnection)

	// make AreYouOpen RPC Call
	doAreYouOpen(serviceClient)
}
