package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

//Power is a control bool to be accessed to shut down the
//clientserver
var Power bool = true
var backendServers map[string]string = make(map[string]string)
var shutdownchan chan string

func main() {

	go StartReverseProxy("8080")

	go GrabServers()
	<-shutdownchan
}

//StartReverseProxy begins the hosting process for the
//client to server application
func StartReverseProxy(port string) {
	fmt.Println("Launching server...")

	ln, _ := net.Listen("tcp", ":"+port)

	fmt.Println("Online - Now Listening On Port: " + port)
	fmt.Println()

	ConnSignal := make(chan string)

	for Power {

		go Session(ln, ConnSignal, port)
		<-ConnSignal

	}
	fmt.Println("Shut Down Signal Sent...Ending")
}

//Session creates a new seesion listening on a port. This
//session handles all interactions with the connected
//client
func Session(ln net.Listener, ConnSignal chan string, port string) {
	conn, _ := ln.Accept()
	ConnSignal <- "New Connection"

	//Checking for server to handle the connecting client
	message, _ := bufio.NewReader(conn).ReadString('\n')
	var serverConn net.Conn = nil
	var err error
	for k, v := range backendServers {
		if strings.Contains(string(message), k) {
			serverConn, err = net.Dial("tcp", v)

			if err != nil {
				fmt.Fprintf(conn, "Server could not find correct route")
				fmt.Fprintf(conn, "No accepting servers"+"\n")
				return
			}
		}
	}

	InboundMessages := make(chan string)
	OutboundMessages := make(chan string)
	go SessionWriter(conn, OutboundMessages)
	go SessionWriter(serverConn, InboundMessages)
	go SessionListener(serverConn, OutboundMessages)
	fmt.Fprintf(serverConn, message)
	SessionListener(conn, InboundMessages)

}

//SessionListener listens for connections noise and sends it to the writer
func SessionListener(Conn1 net.Conn, messages chan string) {
	for {
		message, _ := bufio.NewReader(Conn1).ReadString('\n')
		messages <- (message)
	}

}

//SessionWriter listens for messages channel and sends it to the correct server
func SessionWriter(Conn1 net.Conn, messages chan string) {
	for {
		NewMessage := <-messages
		fmt.Fprintf(Conn1, string(NewMessage))
	}
}

//GrabServers allows user to add servers to list
func GrabServers() {
	for {
		fmt.Println("Grab Servers By Entering in a full address such as Host:Port")

		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.TrimRight(string(text), " \n")
		text = strings.TrimSpace(string(text))
		conn := text

		fmt.Println("What will the server be identified by?")

		reader = bufio.NewReader(os.Stdin)
		text, _ = reader.ReadString('\n')
		text = strings.TrimRight(string(text), " \n")
		text = strings.TrimSpace(string(text))

		servertype := text

		backendServers[servertype] = conn
	}

}
