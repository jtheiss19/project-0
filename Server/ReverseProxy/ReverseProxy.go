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
var logfile *os.File

func main() {
	var err error
	logfile, err = os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println(err)
	}

	go StartReverseProxy("8081")

	go GrabServers()
	<-shutdownchan
}

//StartReverseProxy begins the hosting process for the
//client to server application
func StartReverseProxy(port string) {
	fmt.Println("Launching Reverse Proxy server...")
	logfile.Write([]byte("Launching Reverse Proxy server...\n"))

	ln, _ := net.Listen("tcp", ":"+port)

	fmt.Println("Online - Now Listening On Port: " + port)
	logfile.Write([]byte("Online - Now Listening On Port: " + port + "\n"))

	fmt.Println()

	ConnSignal := make(chan string)

	for Power {

		go Session(ln, ConnSignal, port)
		logfile.Write([]byte(<-ConnSignal))

	}

	logfile.Write([]byte("Shut Down Signal Sent...Ending \n"))

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

	InboundMessages := make(chan string)
	OutboundMessages := make(chan string)
	go SessionWriter(conn, OutboundMessages)
	go SessionWriter(serverConn, InboundMessages)
	go SessionListener(serverConn, OutboundMessages)
	SessionListener(conn, InboundMessages)
}

//SessionListener listens for connections noise and sends it to the writer
func SessionListener(Conn1 net.Conn, messages chan string) {
	for {
		buf := make([]byte, 1024)
		Conn1.SetReadDeadline(time.Now().Add(30 * time.Second))
		_, err := Conn1.Read(buf)
		if err != nil {
			fmt.Println(err)
			Conn1.Write([]byte("Timeout Error, No Signal. Disconnecting. \n"))
			logfile.Write([]byte("Timeout Error, No Signal. Disconnecting. \n"))
			break
		}

		var temp []byte
		for i := 0; i < len(buf); i++ {
			if buf[i] == byte('\u0000') {
				temp = append(buf[0:i])
				break
			}
		}
		logfile.Write([]byte("Recieved message: " + string(temp) + "\n"))

		messages <- string(buf)
	}
}

//SessionWriter listens for messages channel and sends it to the correct server
func SessionWriter(Conn1 net.Conn, messages chan string) {
	for {
		NewMessage := <-messages
		logfile.Write([]byte("Sent message: " + NewMessage + "\n"))
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
		}
	}

}
