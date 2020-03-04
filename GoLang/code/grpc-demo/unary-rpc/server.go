package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	service1 "src/unary-rpc/hello"
)

type server struct {
	service1.GreeterServer
}

func (s *server) SayHello(ctx context.Context, in *service1.HelloRequest) (*service1.HelloReply, error) {
	fmt.Println("getRequest:", in.GetName())
	return &service1.HelloReply{Message: "Hello" + in.GetName()}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *service1.HelloRequest) (*service1.HelloReply, error) {
	return &service1.HelloReply{Message: "Hello Again" + in.GetName()}, nil
}

func main() {
	listen,err := net.Listen("tcp",":1234")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	service1.RegisterGreeterServer(s, &server{})
	if err := s.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
