package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {

	var conn net.Conn
	var err error

	fmt.Println("Enter the address of the server to connect to")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimRight(string(text), " \n")
	text = strings.TrimSpace(string(text))

	for {
		conn, err = net.Dial("tcp", text)

		if err == nil {
			break
		}
	}
	defer conn.Close()
	fmt.Println("Made connection on port 8081")
	conn.Write([]byte("Client"))
	go Writer(conn)
	go Ping(conn)
	Listener(conn)

}

//Listener listens for a server response
func Listener(conn net.Conn) {
	for {
		buf := make([]byte, 1024)

		conn.SetReadDeadline(time.Now().Add(30 * time.Second))
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			conn.Write([]byte("Timeout Error, No Signal. Disconnecting."))
			break
		}
		message := string(buf)
		if strings.Contains(message, string('\u0007')) {
			fmt.Print(message)
			os.Exit(0)
		} else if strings.Contains(message, "ping") {
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

//Ping pings the connection
func Ping(conn net.Conn) {
	for {
		conn.Write([]byte("ping"))
		time.Sleep(20 * time.Second)
	}
}
