package main

import (
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	pb "src/client-streaming-rpc/hello"
	"strings"
)

type server struct {
	pb.GreeterServer
}

func (s *server) SayHello(gs pb.Greeter_SayHelloServer) error {
	var names []string
	for {
		in, err := gs.Recv()
		if err == io.EOF {
			_ = gs.SendAndClose(&pb.HelloReply{Message: "Hello " + strings.Join(names, ",")})
			return nil
		}
		if err != nil {
			log.Printf("failed to recv: %v", err)
			return err
		}
		names = append(names, in.Name)
	}
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
