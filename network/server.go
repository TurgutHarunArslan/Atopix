package network

import (
	"fmt"
	"log"
	"net"
)

func startServer() {
	server, err := net.Listen("tcp4", ":3000")

	if err != nil {
		log.Fatal("Server couldnt be initlized , ", err)
	}
	defer server.Close()

	for {
		client, err := server.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(client)
	}
}
