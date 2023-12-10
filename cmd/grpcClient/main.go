package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/andrewsjuchem/go-expert-grpc/internal/pb"
	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = flag.String("addr", "localhost:50051", "the address to connect to")

func main() {
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()
	categoryServiceClient := pb.NewCategoryServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Specify the fields you want in the response using a FieldMask
	req := &pb.CategoryGetRequest{
		Id: "your-category-id", // Replace with your category ID
		FieldMask: &field_mask.FieldMask{
			Paths: []string{"id", "name"}, // Replace with the fields you want to include
		},
	}

	category, err := categoryServiceClient.GetCategory(ctx, req)
	if err != nil {
		log.Fatalf("Error getting category list: %v", err)
	}
	fmt.Println(category, " ")
}
