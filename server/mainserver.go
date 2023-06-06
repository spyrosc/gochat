package main

import (
	"fmt"
	"net"

	"example.com/server/service"
)

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err.Error())
		return

	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		go func(c net.Conn) {
			defer c.Close()

			closeChannel := make(chan bool)
			fmt.Println("a connection was established")

			go service.WriteToConnection(conn, closeChannel)
			go service.ReadFromConnection(conn)

			<-closeChannel

		}(conn)
	}
}
