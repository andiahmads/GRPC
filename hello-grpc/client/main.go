package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "learning-grpc/hello-grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

// define server
var addr string = "0.0.0.0:2023"

func main() {
	// setup grpc client
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewHelloWorldlServiceClient(conn)

	// call helloWold function
	// HelloWorld(c)

	// call helloworld with streaming server
	// doHelloWorldWithServerStreming(c)

	// call client steaming
	// doHelloWorldWithClientStreaming(c)

	// call bidireactional streaming
	// doHelloWorldWithBIDirectioanl(c)

	// call helloworld with error message
	// doHelloWorldWithErrMsg(c, "")

	// call helloworld with deadline context
	doHelloWorldWithContextDeadlines(c, 1*time.Second)
}

func HelloWorld(c pb.HelloWorldlServiceClient) {

	res, err := c.HelloWorld(context.Background(), &pb.HelloWorldRequest{
		FirstName: "hello-grpc",
		Msisdn:    "082386597687",
	})

	if err != nil {
		log.Fatalf("Could not helloWold:%v\n", err)
	}

	fmt.Println("response: ", res)
}

func doHelloWorldWithServerStreming(c pb.HelloWorldlServiceClient) {
	log.Println("doHelloWorldWithServerStreming was invoked")

	req := &pb.HelloWorldRequest{
		FirstName: "andi",
		Msisdn:    "23234234",
	}

	stream, err := c.HelloWorldWithServerSteaming(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling  doHelloWorldWithServerStreming %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream %v\n", err)
		}

		log.Printf("doHelloWorldWithServerStreming: %v\n", msg.Resutl)
	}
}

// implement Client Streaming
func doHelloWorldWithClientStreaming(c pb.HelloWorldlServiceClient) {
	log.Printf("doHelloWorldWithClientStreaming function was invoked")

	reqs := []*pb.HelloWorldRequest{
		{FirstName: "andi", Msisdn: "123123123"},
		{FirstName: "ahmad", Msisdn: "123123123"},
		{FirstName: "saputra", Msisdn: "123123123"},
		{FirstName: "hahaha", Msisdn: "123123123"},
	}

	stream, err := c.HelloWorldWithClientStreaming(context.Background())

	if err != nil {
		log.Fatalf("Error while calling doHelloWorldWithClientStreaming %v\n", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(10 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receive response from doHelloWorldWithClientStreaming %v\n", err)
	}

	log.Printf("HelloWorldClientStreaming response: %s\n", res.Resutl)

}

// implement bi directioanl grpc
func doHelloWorldWithBIDirectioanl(c pb.HelloWorldlServiceClient) {

	fmt.Println("doHelloWorldWithBIDirectioanl function was invoked")

	stream, err := c.HelloWorldWithBidirectional(context.Background())

	if err != nil {
		log.Fatalf("Error calling doHelloWorldWithBIDirectioanl %v\n", err)
	}

	reqs := []*pb.HelloWorldRequest{
		{FirstName: "andi", Msisdn: "123123123"},
		{FirstName: "ahmad", Msisdn: "123123123"},
		{FirstName: "saputra", Msisdn: "123123123"},
		{FirstName: "hahaha", Msisdn: "123123123"},
	}

	wait := make(chan struct{})

	go func() {

		for _, req := range reqs {
			log.Printf("Send Request: %v\n", req)

			stream.Send(req)
			time.Sleep(3 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving: %v\n", err)

			}

			log.Printf("Received: %v\n", res.Resutl)
		}

		close(wait)

	}()

	<-wait
}

func doHelloWorldWithErrMsg(c pb.HelloWorldlServiceClient, firstName string) {
	log.Printf("doHelloWorldWithErrMsg was invoked")

	res, err := c.HelloWorldWithErrorMsg(
		context.Background(),
		&pb.HelloWorldRequest{FirstName: firstName, Msisdn: "23478273487"},
	)

	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Printf("Error message from server: %s\n", e.Message())
			log.Printf("Error code from server: %s\n", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Printf("we probbaly send nul value")
				return
			}
		} else {
			log.Fatalf("A no GRPC err: %v\n", err)
		}
	}
	log.Printf("received message from helloworld: %s\n", res.Resutl)
}

func doHelloWorldWithContextDeadlines(c pb.HelloWorldlServiceClient, timeout time.Duration) {
	log.Println("doHelloWorldWithContextDeadlines was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.HelloWorldRequest{
		FirstName: "andi ahamd",
		Msisdn:    "9090909234234",
	}

	res, err := c.HelloWorldWithDeadLines(ctx, req)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Printf("Deadline exceeded!")
				return
			} else {
				log.Fatalf("Unexpected GRPC error: %v\n", err)
			}
		} else {
			log.Fatalf("A non GRPC error:%v\n", err)
		}
	}

	log.Printf("receiving data: %s\n", res.Resutl)

}
