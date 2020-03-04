package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "src/bidirectional-streaming-rpc/hello"
	"strconv"
	"sync"
)

func Send(wg *sync.WaitGroup) {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
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
		err = stream.Send(&pb.HelloRequest{Name: "world" + strconv.Itoa(i)})
		if err != nil {
			log.Printf("failed to send: %v", err)
		}
		reply, err := stream.Recv()
		if err != nil {
			log.Printf("failed to recv: %v", err)
			break
		}
		log.Printf("Greeting: %s", reply.Message)
	}
	_ = stream.CloseSend()
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go Send(wg)
	}
	wg.Wait()
	//conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer conn.Close()
	//c := pb.NewGreeterClient(conn)
	//stream, err := c.SayHello(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for i := 0; i < 100; i++ {
	//	err = stream.Send(&pb.HelloRequest{Name: "world" + strconv.Itoa(i)})
	//	if err != nil {
	//		log.Printf("failed to send: %v", err)
	//	}
	//	reply, err := stream.Recv()
	//	if err != nil {
	//		log.Printf("failed to recv: %v", err)
	//		break
	//	}
	//	log.Printf("Greeting: %s", reply.Message)
	//}
	//_ = stream.CloseSend()
}
