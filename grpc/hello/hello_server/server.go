package main

import (
	"fmt"
	"github.com/younghyun-ahn/go-etc/grpc/hello/hellopb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func (*server) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	fmt.Printf("Hello function was invoked with %v", req)
	firstName := req.GetHello().GetFirstName()
	result := "Hello " + firstName
	res := &hellopb.HelloResponse{
		Result: result,
	}
	return res, nil
}

//Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)

func main() {
	fmt.Println("Hello server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Faild to listen: %v", err)
	}

	s := grpc.NewServer()
	hellopb.RegisterHelloServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
