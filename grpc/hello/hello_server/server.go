package main

import (
	"fmt"
	"github.com/younghyun-ahn/go-etc/grpc/hello/hellopb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"time"
)

type server struct{}

func (*server) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	fmt.Printf("Hello function was invoked with %v\n", req)
	firstName := req.GetHello().GetFirstName()
	result := "Hello " + firstName
	res := &hellopb.HelloResponse{
		Result: result,
	}
	return res, nil
}

func (*server) HelloWithDeadline(ctx context.Context, req *hellopb.HelloWithDeadlineRequest) (*hellopb.HelloWithDeadlineResponse, error) {
	fmt.Printf("HelloWithDeadline function was invoked with %v\n", req)
	for i := 0; i < 3; i ++ {
		if ctx.Err() == context.Canceled {
			// the client canceled the request
			fmt.Println("The client canceled the request!")
			return nil, status.Error(codes.Canceled, "The client canceled the request")
		}
		time.Sleep(1* time.Second)
	}
	firstName := req.GetHello().GetFirstName()
	result := "Hello " + firstName
	res := &hellopb.HelloWithDeadlineResponse{
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
