package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/andrewsjuchem/go-expert-grpc/internal/database"
	"github.com/andrewsjuchem/go-expert-grpc/internal/pb"
	"github.com/andrewsjuchem/go-expert-grpc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Connect to the database
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	grpcServer := grpc.NewServer()

	// Category
	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)

	// Course
	courseDb := database.NewCourse(db)
	courseService := service.NewCourseService(*courseDb)
	pb.RegisterCourseServiceServer(grpcServer, courseService)

	reflection.Register(grpcServer)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
