package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	donutspb "training/donuts/protos"
	"training/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

var collection *mongo.Collection

type server struct {
	donutspb.UnimplementedDonutsServiceServer
}

type donutItem struct {
	ID     primitive.ObjectID `bson: "_id,omitempty"`
	Name   string             `bson: "name"`
	Price  int32              `bson: "price"`
	Rating float32            `bson: "rating"`
	Count  int32              `bson: "count"`
}

// RPC to serve AreYouOpen requests
func (server) AreYouOpen(ctx context.Context, req *donutspb.AreYouOpenRequest) (*donutspb.AreYouOpenResponse, error) {
	fmt.Println("received AreYouOpen RPC Call...")
	// TODO: Use this list to validate input
	// days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	openDays := []string{"Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	dayOfTheWeek := req.GetDayOfTheWeek()

	isPresent, _ := utils.Contains(openDays, dayOfTheWeek)

	/*
		If a scalar message field is set to its default, the value will not be serialized on the wire.
		Because of this the response is empty object if the store is closed.
	*/
	return &donutspb.AreYouOpenResponse{
		Reponse: isPresent,
	}, nil
}

// RPC to server CreateDonuts requests
func (server) CreateDonuts(ctx context.Context, req *donutspb.CreateDonutsRequest) (*donutspb.CreateDonutsResponse, error) {

	// Read request data
	donut := req.GetDonut()
	count := req.GetCount()

	// Create data object to be inserted into DB
	data := donutItem{
		Name:   donut.GetName(),
		Price:  donut.GetPrice(),
		Rating: donut.GetRating(),
		Count:  count,
	}

	insertResponse, insertError := collection.InsertOne(context.Background(), data)

	// Handle insertion error
	if insertError != nil {
		fmt.Printf("\n Error inserting donuts %v \n", insertError)
		return nil, status.Errorf(codes.Internal, "Error creating new donut record")
	}

	oid, ok := insertResponse.InsertedID.(primitive.ObjectID)

	// Handle error while reading the ID
	if !ok {
		return nil, status.Errorf(codes.Internal, "Error converting to OID")
	}

	fmt.Printf("\nCreated new donut record %v", oid)

	//TODO: fill in the response
	return &donutspb.CreateDonutsResponse{
		Created: true,
	}, nil
}

// RPC to serve GetDonutsList request
func (server) GetDonutsList(ctx context.Context, req *donutspb.GetDonutsListRequest) (*donutspb.GetDonutsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDonutsList not implemented")
}

// RPC to serve OrderDonuts request
func (server) OrderDonuts(ctx context.Context, req *donutspb.OrderDonutsRequest) (*donutspb.OrderDonutsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderDonuts not implemented")
}

// RPC to serve UpdateDonuts request
func (server) UpdateDonuts(ctx context.Context, req *donutspb.UpdateDonutsRequest) (*donutspb.UpdateDonutsReponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDonuts not implemented")
}

// RPC to serve ClearDonuts request
func (server) ClearDonuts(ctx context.Context, req *donutspb.ClearDonutsRequest) (*donutspb.ClearDonutsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearDonuts not implemented")
}

func main() {

	fmt.Println("starting donuts db...")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	mongoClient, dbConnectError := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if dbConnectError != nil {

		log.Fatalf("Error connecting to database %v", dbConnectError)
	}

	collection = mongoClient.Database("donutsdb").Collection("donuts")

	fmt.Println("starting donuts server...")

	// create listener
	listener, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()
	donutsServer := &server{}

	donutspb.RegisterDonutsServiceServer(s, donutsServer)

	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve %v", err)
		}
	}()

	// Register reflection service on gRPC server.
	reflection.Register(s)

	//Wait for Ctrl+C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch

	fmt.Println("stopping the server...")
	s.Stop()

	fmt.Println("closing the listener...")
	listener.Close()

	fmt.Println("disconnecting mongodb client...")
	mongoClient.Disconnect(ctx)

	fmt.Println("good bye...")
}
