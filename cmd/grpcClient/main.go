package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/andrewsjuchem/go-expert-grpc/internal/pb"
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
	fieldMask := &pb.FieldMask{
		IncludedFields: []string{"name"},
	}

	categoryList, err := categoryServiceClient.ListCategories(ctx, fieldMask)
	if err != nil {
		log.Fatalf("Error getting category list: %v", err)
	}
	for _, value := range categoryList.Categories {
		fmt.Println(value, " ")
	}
}
