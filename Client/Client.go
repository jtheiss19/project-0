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
	defer conn.Close()
	fmt.Println("Made connection on port 8080")
	conn.Write([]byte("Client"))
	go Writer(conn)
	Listener(conn)

}

//Listener listens for a server response
func Listener(conn net.Conn) {
	for {
		buf := make([]byte, 2048)

		conn.Read(buf)
		message := string(buf)
		if strings.Contains(message, string('\u0007')) {
			fmt.Print(message)
			os.Exit(0)
		} else {
			fmt.Print(message)
		}
	}
}

//Writer writes to the server
func Writer(conn net.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		bytes, _ := reader.ReadBytes('\n')
		bytes = append(bytes[:len(bytes)-1])
		conn.Write([]byte(bytes))
	}
}
