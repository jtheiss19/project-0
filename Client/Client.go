package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	var conn net.Conn
	var err error

	for {
		conn, err = net.Dial("tcp", "127.0.0.1:8080")

		if err == nil {
			break
		}
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Command to run: ")
		text, _ := reader.ReadString('\n')

		fmt.Fprintf(conn, text+"\n")

		if text == "Disconnect \n" || text == "Power \n" {
			break
		}

		message, _ := bufio.NewReader(conn).ReadString('\u0000')
		fmt.Print("Message from server: " + message)
	}
}
