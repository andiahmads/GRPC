package main

import (
	"context"
	"fmt"
	"io"
	pb "learning-grpc/hello-grpc/proto"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

var addr string = "0.0.0.0:2023"

type Server struct {
	pb.HelloWorldlServiceServer
}

func main() {
	// make listen serv
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on :%v\n", err)
	}

	log.Printf("Listening to serve: %v\n", addr)

	// implement grpc server
	s := grpc.NewServer()

	// register hello world server
	pb.RegisterHelloWorldlServiceServer(s, &Server{})

	reflection.Register(s)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}

}

// implement GRPC server with UNARY GRPC
func (s *Server) HelloWorld(
	ctx context.Context,
	in *pb.HelloWorldRequest,
) (*pb.HelloWorldResponse, error) {
	fmt.Printf("Hello World function was invoked with %v\n", in)
	userCreatedDate := time.Now().Format("2006-01-02 15:04:05")

	return &pb.HelloWorldResponse{
		Resutl: in.FirstName + " " + userCreatedDate,
		Msisdn: in.Msisdn,
	}, nil

}

// implement Server Steaming
func (s *Server) HelloWorldWithServerSteaming(
	req *pb.HelloWorldRequest,
	stream pb.HelloWorldlService_HelloWorldWithServerSteamingServer,
) error {
	fmt.Printf("HelloWorldWithServerStreaming service was invoked %v\n", req)

	for i := 0; i < 100; i++ {
		res := fmt.Sprintf("hello %s, with msisdn %s, number %d", req.FirstName, req.Msisdn, i)

		stream.Send(&pb.HelloWorldResponse{
			Resutl: res,
		})
	}
	return nil
}

// implement Client Streaming
func (s *Server) HelloWorldWithClientStreaming(
	stream pb.HelloWorldlService_HelloWorldWithClientStreamingServer,
) error {
	log.Printf("HelloWorldWithClientStreaming function was invoked")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.HelloWorldResponse{
				Resutl: res,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		log.Printf("Receiving: %v\n", req)

		res += fmt.Sprintf("Hello %s!\n your number is %s\n", req.FirstName, req.Msisdn)
	}
}

// implement Bidirectional grpc
func (s *Server) HelloWorldWithBidirectional(
	stream pb.HelloWorldlService_HelloWorldWithBidirectionalServer,
) error {
	log.Println("Helloworld with Bidirectional function was invoked")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("error while reading client stream %v\n", err)
		}

		res := "hello " + req.FirstName + "wiht msisdn " + req.Msisdn
		err = stream.Send(&pb.HelloWorldResponse{
			Resutl: res,
		})
		log.Printf("Receiving: %v\n", req)
		if err != nil {
			log.Fatalf("Error while sending data to client %v\n", err)
		}
	}

}

func (s *Server) HelloWorldWithErrorMsg(
	ctx context.Context,
	req *pb.HelloWorldRequest,
) (*pb.HelloWorldResponse, error) {
	log.Printf("Helloworld with error message function was invoked: %v\n", req)

	if req.FirstName == "" {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("received null value: %s", req.FirstName),
		)
	}

	return &pb.HelloWorldResponse{
		Resutl: req.FirstName + " " + req.Msisdn,
	}, nil
}

func (s *Server) HelloWorldWithDeadLines(
	ctx context.Context,
	req *pb.HelloWorldRequest,
) (*pb.HelloWorldResponse, error) {
	log.Printf("HelloWorldWithDeadLines function was invoked with %v\n", req)
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("the client canceled the request!!")
			return nil, status.Error(codes.Canceled, "the client canceled the request!")
		}

		time.Sleep(3 * time.Second)
	}
	return &pb.HelloWorldResponse{
		Resutl: req.FirstName + " " + req.Msisdn,
	}, nil
}
