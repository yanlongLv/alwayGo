package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"github.com/alwayGo/newProject/api/helloworld"
)

func main() {
	conn, err := grpc.Dial(":8972", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("faild to connect: %v", err)
	}
	defer conn.Close()
	c := helloworld.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &helloworld.HelloRequest{Name: "q1mi"})
	if err != nil {
		fmt.Printf("could not greet: %v", err)
	}
	fmt.Printf("Greeting: %s !\n", r.Message)
}
