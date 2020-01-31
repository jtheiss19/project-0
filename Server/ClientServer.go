package server

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

//Power is a control bool to be accessed to shut down the
//clientserver
var Power bool = true
var connections []net.Conn
var logfile *os.File

//StartClientServer begins the hosting process for the
//client to server application
func StartClientServer(port string) {
	var err error
	logfile, err = os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Launching Client server...")
	Write([]byte("Launching Client server..."), "N/A", "N/A", port)

	ln, _ := net.Listen("tcp", ":"+port)

	fmt.Println("Online - Now Listening On Port: " + port)
	Write([]byte("Online - Now Listening On Port: "+port), "N/A", "N/A", port)

	fmt.Println()

	ConnSignal := make(chan string)

	for Power {

		go Session(ln, ConnSignal, port)
		<-ConnSignal

	}
	fmt.Println("Shutting Down...")
	Write([]byte("Shutting Down..."), "N/A", "N/A", port)

	for i := 0; i < len(connections); i++ {
		connections[i].Write([]byte("Server is shutting down, Disconnecting you" + string('\u0007') + "\n"))
		Write([]byte("Server is shutting down, Disconnecting you"+string('\u0007')), connections[i].LocalAddr().String(), connections[i].RemoteAddr().String(), connections[i].LocalAddr().String())
	}
	Write([]byte("Shut Down Signal Sent...Ending"), "N/A", "N/A", port)
}

//Session creates a new seesion listening on a port. This
//session handles all interactions with the connected
//client
func Session(ln net.Listener, ConnSignal chan string, port string) {
	conn, _ := ln.Accept()
	defer conn.Close()
	connections = append(connections, conn)

	fmt.Println("New Connection On")
	Write([]byte("New Connection"), conn.LocalAddr().String(), "N/A", conn.LocalAddr().String())
	ConnSignal <- "New Connection"

	messages := make(chan []string)
	go SessionWriter(messages, conn)
	go Ping(conn)
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
		Write([]byte(out), conn.LocalAddr().String(), conn.RemoteAddr().String(), conn.LocalAddr().String())
		conn.Write(out)
	}
}

//SessionListener handles all incoming messages from a client
//and parses them for commands before passing it to the writer
func SessionListener(messages chan []string, conn net.Conn, ConnSignal chan string) {
	for {
		buf := make([]byte, 1024)

		conn.SetReadDeadline(time.Now().Add(30 * time.Second))
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			conn.Write([]byte("Timeout Error, No Signal. Disconnecting. \n" + string('\u0007')))
			Write([]byte("Timeout Error, No Signal. Disconnecting."), conn.LocalAddr().String(), conn.RemoteAddr().String(), conn.LocalAddr().String())
			break
		}

		//CRITICAL: removes null bytes from buffer
		for i := 0; i < len(buf); i++ {
			if buf[i] == byte('\u0000') {
				buf = append(buf[0:i])
				break
			}
		}

		command := strings.TrimSpace(string(buf))
		commandSlice := strings.Split(command, " ")

		Write(buf, conn.RemoteAddr().String(), conn.LocalAddr().String(), conn.LocalAddr().String())

		switch string(buf) {

		case "ping":

		case "":
			fmt.Println()
			buf = []byte("Disconnect")

		case "Power":
			Power = false
			conn.Write([]byte("Server is shutting down. Disconnecting you \n" + string('\u0007')))
			ConnSignal <- "Remoted Power Toggled from "
			return

		case "Disconnect":
			conn.Write([]byte("Disconnecting you from Server \n" + string('\u0007') + string("\n")))
			conn.Close()

			return
		}

		messages <- commandSlice

	}
}

//Ping pings the connection
func Ping(conn net.Conn) {
	for {
		conn.Write([]byte("ping"))
		time.Sleep(20 * time.Second)
	}
}

var mu = &sync.Mutex{}

//Write writes to a logfile
func Write(info []byte, In string, Out string, WhoAmI string) {
	//Build complete String
	FullString := WhoAmI + ", " + In + ", " + Out + ", " + string(info) + "\n"
	mu.Lock()
	defer mu.Unlock()
	logfile.Write([]byte(FullString))
}
