package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
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
	logfile.Write([]byte("Launching Client server...\n"))

	ln, _ := net.Listen("tcp", ":"+port)

	fmt.Println("Online - Now Listening On Port: " + port)
	logfile.Write([]byte("Online - Now Listening On Port: " + port + "\n"))

	fmt.Println()

	ConnSignal := make(chan string)

	for Power {

		go Session(ln, ConnSignal, port)
		logfile.Write([]byte(<-ConnSignal))

	}
	fmt.Println("Shutting Down...")
	logfile.Write([]byte("Shutting Down...\n"))

	for i := 0; i < len(connections); i++ {
		connections[i].Write([]byte("Server is shutting down, Disconnecting you" + string('\u0007') + "\n"))
	}
	logfile.Write([]byte("Shut Down Signal Sent...Ending"))
}

//Session creates a new seesion listening on a port. This
//session handles all interactions with the connected
//client
func Session(ln net.Listener, ConnSignal chan string, port string) {
	conn, _ := ln.Accept()
	defer conn.Close()
	connections = append(connections, conn)

	fmt.Println("New Connection On Port: " + port)
	fmt.Println()
	ConnSignal <- "New Connection On Port: " + port + "\n"

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
		logfile.Write([]byte("Sending Command: " + string(out) + "\n"))
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

		conn.SetReadDeadline(time.Now().Add(30 * time.Second))
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			conn.Write([]byte("Timeout Error, No Signal. Disconnecting. \n" + string('\u0007')))
			logfile.Write([]byte("Timeout Error, No Signal. Disconnecting. \n" + string('\u0007')))
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

		logfile.Write([]byte("Raw Text Received From " + host + ": " + string(buf)))

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
		conn.Write([]byte("ping\n"))
		time.Sleep(20 * time.Second)
	}
}
