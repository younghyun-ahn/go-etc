package main

import (
	"fmt"
	"github.com/younghyun-ahn/go-etc/grpc/hello/hellopb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
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

	doUnary(c)
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
