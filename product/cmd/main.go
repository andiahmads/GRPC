package main

import (
	"learning-grpc/product/config"
	"learning-grpc/product/services"
	"log"
	"net"

	productPb "learning-grpc/product/proto"

	"google.golang.org/grpc"
)

const (
	port = ":2022"
)

func main() {
	netListen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen %v\n", err.Error())
	}

	db := config.ConnectDataBase()

	grpcServer := grpc.NewServer()

	productService := services.ProductService{DB: db}

	productPb.RegisterProductServiceServer(grpcServer, &productService)

	log.Printf("server started at %v\n", netListen.Addr())

	if err := grpcServer.Serve(netListen); err != nil {
		log.Fatalf("Failed to serve %v\n", err.Error())
	}

}
