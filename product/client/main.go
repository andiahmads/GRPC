package main

import (
	"context"
	"fmt"
	"log"

	productPb "learning-grpc/product/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
)

var addr string = "0.0.0.0:2022"

func main() {
	// setup grpc client
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := productPb.NewProductServiceClient(conn)
	getProducts(c)

}

func getProducts(c productPb.ProductServiceClient) {

	res, err := c.GetProducts(context.Background(), &productPb.Page{
		Page: proto.Int64(2),
	})

	if err != nil {
		log.Fatalf("could not get data product%v\n", err.Error())
	}

	data, err := proto.Marshal(res)
	if err != nil {
		log.Fatalf("marshall error %v\n", err)
	}

	testProducts := &productPb.Products{}
	if err := proto.Unmarshal(data, testProducts); err != nil {
		log.Fatal("Unmarshal error", err)
	}

	fmt.Println("response: ", testProducts)

}

func test() {

}
