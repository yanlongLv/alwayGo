package main

import (
	"context"
	"log"
	"net"

	grpc "google.golang.org/grpc"

	pb "github.com/alwayGo/newProject/api/helloworld"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "hello " + in.GetName()}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8972")
	if err != nil {
		log.Fatal("failed error ", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err = s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
