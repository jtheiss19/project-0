package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
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

//Listener listens for a server response
func Listener(conn net.Conn) {
	for {
		message, _ := bufio.NewReader(conn).ReadString('\u0000')
		if strings.Contains(message, string('\u0007')) {
			fmt.Print(message)
			os.Exit(0)
		}
		fmt.Print(message)
	}
}

//Writer writes to the server
func Writer(conn net.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		fmt.Fprintf(conn, text+"\n")
	}
}
