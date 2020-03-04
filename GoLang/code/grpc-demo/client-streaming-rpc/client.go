package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	pb "src/client-streaming-rpc/hello"
	"strconv"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	stream, err := c.SayHello(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 100; i++ {
		err := stream.Send(&pb.HelloRequest{Name: "world" + strconv.Itoa(i)})
		if err != nil {
			log.Printf("failed to call: %v", err)
			break
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Printf("failed to recv: %v", err)
	}
	log.Printf("Greeting: %s", reply.Message)
}
