package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

const max = 10

func main() {
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenErr=", err)
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal("connErr=", err)
		}
		go handlerClient(conn)
	}
}

func handlerClient(conn net.Conn) {
	defer conn.Close()
	_ = conn.SetDeadline(time.Now().Add(1 * time.Minute))
	info := make([]byte, max)
	for {
		n, err := conn.Read(info)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("收到客户端消息:", string(info[:n]))
		reply(conn, []byte("recv data"))
	}
}

func reply(conn net.Conn, msg []byte) {
	_, err := conn.Write(msg)
	if err != nil {
		log.Fatal("replyError=", err)
	}
}
