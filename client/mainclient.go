package main

import (
	"fmt"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err.Error())
	}
	for {
		readB := make([]byte, 1024)
		n, _ := conn.Read(readB)

		if n > 0 {
			fmt.Println(string(readB))
		}
	}

}
