package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	pb "src/server-streaming-rpc/hello"
	"strconv"
)

type server struct {
	pb.GreeterServer
}

func (s *server) SayHello(in *pb.HelloRequest, gs pb.Greeter_SayHelloServer) error {
	name := in.Name
	for i := 0; i < 100; i++ {
		err := gs.Send(&pb.HelloReply{Message: "hello" + name + strconv.Itoa(i)})
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
