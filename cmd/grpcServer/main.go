package main

import (
	"database/sql"
	"fmt"
	"net"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rcarvalho-pb/go-grpc/internal/database"
	"github.com/rcarvalho-pb/go-grpc/internal/pb"
	"github.com/rcarvalho-pb/go-grpc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}

	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	port := ":50051"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	fmt.Println("Listening on port", port)

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
