package main

import (
	"fmt"
	"github.com/younghyun-ahn/go-etc/grpc/hello/hellopb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

func main() {
	fmt.Println("Hello client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	c := hellopb.NewHelloServiceClient(conn)
	//fmt.Printf("Created client: %f", c)

	//doUnary(c)
	doDeadline(c, 5*time.Second) //should complete
	doDeadline(c, 1*time.Second) //should timeout
}

func doUnary(c hellopb.HelloServiceClient) {
	req := &hellopb.HelloRequest{
		Hello: &hellopb.Hello{
			FirstName: "Younghyun",
			LastName:  "Ahn",
		},
	}
	res, err := c.Hello(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Hello RPC: %v", err)
	}
	log.Printf("Response from Hello: %v", res.Result)
}

func doDeadline(c hellopb.HelloServiceClient, timeout time.Duration) {
	req := &hellopb.HelloWithDeadlineRequest{
		Hello: &hellopb.Hello{
			FirstName: "Younghyun",
			LastName: "Ahn",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	res, err := c.HelloWithDeadline(ctx, req)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Timeout was hit! Deadline was exceeded")
			} else {
				fmt.Printf("Unexpected error: %v", statusErr)
			}
		} else {
			log.Fatalf("Error while calling Hello RPC: %v", err)
		}
		return
	}
	log.Printf("Response from Hello: %v", res.Result)

}
