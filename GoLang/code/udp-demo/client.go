package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	var input []byte
	var recv = make([]byte, 100)
	for {
		fmt.Print("input:")
		_, _ = fmt.Scan(&input)
		if string(input) == "stop" {
			break
		}
		n, err := conn.Write(input)
		if err != nil {
			log.Fatal(n, err)
		}
		n, err = conn.Read(recv)
		if err != nil {
			log.Fatal(n, err)
		}
		fmt.Println("收到客户端回复:", string(recv[:n]))
	}
}
