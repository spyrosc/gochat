package service

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func WriteToConnection(c net.Conn, closeChan chan bool) {

	for {
		reader := bufio.NewReader(os.Stdin)

		writeBuf, _, _ := reader.ReadLine()

		if strings.Contains(string(writeBuf), "exit") {
			closeChan <- true
			break
		}

		c.Write(writeBuf[:])

	}

}

func ReadFromConnection(c net.Conn) {

	readBuf := make([]byte, 2048)

	for {

		n, _ := c.Read(readBuf)
		if n > 0 {
			fmt.Println(string(readBuf[:]))

		}
	}

}
