package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	pb "src/server-streaming-rpc/hello"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	stream, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "world"})
	if err != nil {
		log.Fatal(err)
	}
	for {
		reply, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("failed to recv: %v", err)
		}
		log.Printf("Greeting: %s", reply.Message)
	}
}
