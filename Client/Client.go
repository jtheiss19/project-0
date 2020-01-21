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
	go Writer(conn)
	Listener(conn)
}

func Listener(conn net.Conn) {
	for {
		message, _ := bufio.NewReader(conn).ReadString('\u0000')
		if message == "END"+string('\u0000') {
			fmt.Println("Server Has Shut Down, Ending Connection")
			break
		}
		fmt.Print(message)
	}
}

func Writer(conn net.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		fmt.Fprintf(conn, text+"\n")
	}
}
