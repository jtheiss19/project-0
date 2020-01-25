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
	fmt.Fprintf(conn, "Client"+"\n")
	go Writer(conn)
	Listener(conn)

}

//Listener listens for a server response
func Listener(conn net.Conn) {
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		if strings.Contains(message, string('\u0007')) {
			fmt.Print(message)
			os.Exit(0)
		}

		if strings.Contains(message, string('\u0000')) {
			stringArray := strings.Split(message, string('\u0000'))
			for i := 0; i < len(stringArray); i++ {
				fmt.Println(stringArray[i])
			}
		} else {
			fmt.Println(message)
		}
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
