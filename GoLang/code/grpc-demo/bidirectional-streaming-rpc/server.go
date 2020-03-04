package main

import (
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	pb "src/bidirectional-streaming-rpc/hello"
)

type server struct {
	pb.GreeterServer
}

func (s *server) SayHello(gs pb.Greeter_SayHelloServer) error {
	for {
		in, err := gs.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			//log.Fatalf("failed to recv: %v", err)
			fmt.Println(err)
			//无法退出
			return nil
		}
		//fmt.Println("revcive=", in.Name)
		err = gs.Send(&pb.HelloReply{Message: "Hello " + in.Name})
		if err != nil {
			return err
		}
	}
}

func main() {
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()
	s := grpc.NewServer()
	defer s.Stop()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
