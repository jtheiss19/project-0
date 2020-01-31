package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

//Power is a control bool to be accessed to shut down the
//clientserver
var Power bool = true
var backendServers map[string]string = make(map[string]string)
var shutdownchan chan string
var logConn net.Conn

func main() {
	fmt.Println("Reverse Proxy Server Terminal")
	go GrabServers()
	<-shutdownchan
}

//StartReverseProxy begins the hosting process for the
//client to server application
func StartReverseProxy(port string) {
	fmt.Println("Launching Reverse Proxy server...")

	ln, _ := net.Listen("tcp", ":"+port)

	fmt.Println("Online - Now Listening On Port: " + port)

	fmt.Println()

	ConnSignal := make(chan string)

	for Power {

		go Session(ln, ConnSignal, port)
		<-ConnSignal

	}

}

//Session creates a new seesion listening on a port. This
//session handles all interactions with the connected
//client
func Session(ln net.Listener, ConnSignal chan string, port string) {
	conn, _ := ln.Accept()
	defer conn.Close()
	ConnSignal <- "New Connection\n"

	//Checking for server to handle the connecting client
	buf := make([]byte, 1024)
	conn.Read(buf)
	var serverConn net.Conn = nil
	var err error
	for {
		for k, v := range backendServers {
			if strings.Contains(string(buf), k) {
				serverConn, err = net.Dial("tcp", v)
				defer serverConn.Close()
				serverConn.Write(buf)
				if err != nil {
				} else {
					break
				}
			}
		}
		if serverConn != nil {
			break
		}
	}

	Shutdown := make(chan string)
	InboundMessages := make(chan string)
	OutboundMessages := make(chan string)
	go SessionWriter(conn, OutboundMessages, Shutdown)
	go SessionWriter(serverConn, InboundMessages, Shutdown)
	go SessionListener(serverConn, OutboundMessages, Shutdown)
	go SessionListener(conn, InboundMessages, Shutdown)
	<-Shutdown
}

//SessionListener listens for connections noise and sends it to the writer
func SessionListener(Conn1 net.Conn, messages chan string, shutdown chan string) {
	for {
		buf := make([]byte, 1024)
		Conn1.SetReadDeadline(time.Now().Add(30 * time.Second))
		_, err := Conn1.Read(buf)
		if err != nil {
			fmt.Println(err)
			Conn1.Write([]byte("Timeout Error, No Signal. Disconnecting. \n"))
			break
		}

		if logConn != nil {
			logConn.Write(buf)
		}

		messages <- string(buf)
	}
	shutdown <- "Session Closed"
}

//SessionWriter listens for messages channel and sends it to the correct server
func SessionWriter(Conn1 net.Conn, messages chan string, shutdown chan string) {
	for {
		NewMessage := <-messages

		if logConn != nil {
			logConn.Write([]byte(NewMessage))
		}

		Conn1.Write([]byte(NewMessage))
	}
}

//GrabServers allows user to add servers to list
func GrabServers() {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.TrimRight(string(text), " \n")
		text = strings.TrimSpace(string(text))

		switch text {
		case "add":
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

		case "Log":
			fmt.Println("Grab logging server By Entering in a full address such as Host:Port")

			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			text = strings.TrimRight(string(text), " \n")
			text = strings.TrimSpace(string(text))

			conn, _ := net.Dial("tcp", text)

			logConn = conn

		case "Launch":
			fmt.Println("Enter in a port")

			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			text = strings.TrimRight(string(text), " \n")
			text = strings.TrimSpace(string(text))

			go StartReverseProxy(text)

		}
	}

}
