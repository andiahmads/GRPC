package main

import (
	"fmt"
	pb "learning-grpc/product/proto"
	"log"

	"google.golang.org/protobuf/proto"
)

func main() {
	products := &pb.Products{
		Data: []*pb.Product{
			{
				Id:    1,
				Name:  "Nike Black",
				Price: 10000.0,
				Stock: 100,
				Category: &pb.Category{
					Id:   1,
					Name: "shirt",
				},
			},
			{
				Id:    2,
				Name:  "Vans orange",
				Price: 10000.0,
				Stock: 100,
				Category: &pb.Category{
					Id:   2,
					Name: "shoe",
				},
			},
		},
	}

	data, err := proto.Marshal(products)
	if err != nil {
		log.Fatalf("marshall error %v\n", err)
	}

	fmt.Println(data)

	testProducts := &pb.Products{}
	if err := proto.Unmarshal(data, testProducts); err != nil {
		log.Fatal("Unmarshal error", err)
	}

	fmt.Println(testProducts)

}
