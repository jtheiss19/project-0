package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
)

//Power is a control bool to be accessed to shut down the
//clientserver
var Power bool = true
var connections []net.Conn

//StartClientServer begins the hosting process for the
//client to server application
func StartClientServer(port string) {
	fmt.Println("Launching server...")

	ln, _ := net.Listen("tcp", ":"+port)

	fmt.Println("Online - Now Listening On Port: " + port)
	fmt.Println()

	ConnSignal := make(chan string)

	for Power {

		go Session(ln, ConnSignal, port)
		<-ConnSignal

	}
	fmt.Println("Shutting Down...")
	for i := 0; i < len(connections); i++ {
		fmt.Fprintf(connections[i], "Server is shutting down, Disconnecting you"+string('\u0007')+"\n")
	}
	fmt.Println("Shut Down Signal Sent...Ending")
}

//Session creates a new seesion listening on a port. This
//session handles all interactions with the connected
//client
func Session(ln net.Listener, ConnSignal chan string, port string) {
	conn, _ := ln.Accept()
	connections = append(connections, conn)

	fmt.Println("New Connection On Port: " + port)
	fmt.Println()
	ConnSignal <- "New Connection On Port: " + port

	messages := make(chan []string)
	go SessionWriter(messages, conn)
	SessionListener(messages, conn, ConnSignal)

}

//SessionWriter handles all out going and command communication
//with a client
func SessionWriter(messages chan []string, conn net.Conn) {

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		commandSlice := <-messages
		out, err := exec.Command(dir+"/main", commandSlice...).Output()
		if err != nil {
			fmt.Println(err)
			out = []byte("Command is not valid\n")
		}
		conn.Write(out)
	}
}

//SessionListener handles all incoming messages from a client
//and parses them for commands before passing it to the writer
func SessionListener(messages chan []string, conn net.Conn, ConnSignal chan string) {
	host, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	for {
		buf := make([]byte, 1024)

		conn.Read(buf)

		//CRITICAL: removes null bytes from buffer
		for i := 0; i < len(buf); i++ {
			if buf[i] == byte('\u0000') {
				buf = append(buf[0:i])
				break
			}
		}

		command := strings.TrimSpace(string(buf))
		commandSlice := strings.Split(command, " ")

		fmt.Println("Raw Text Received From "+host+": ", string(buf))
		if string(buf) == "" {
			fmt.Println()
			buf = []byte("Disconnect")
		}

		if string(buf) == "Power" {
			Power = false
			ConnSignal <- "Remoted Power Toggled from "
			return
		}

		if string(buf) == "Disconnect" {
			fmt.Fprintf(conn, "Disconnecting you from Server"+string('\u0007')+string("\n"))
			for i := 0; i < len(connections); i++ {
				if connections[i] == conn {
					connections[i] = connections[len(connections)-1]
					connections[len(connections)-1] = nil
					connections = connections[:len(connections)-1]
					fmt.Println("Connection with " + host + " has ended remotely")
					fmt.Println()
				}
			}

			return
		}

		fmt.Println("Command to exectute for "+host+": ", commandSlice)
		fmt.Println()

		messages <- commandSlice

	}
}
