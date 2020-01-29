package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sort"
	"strings"
	"time"
)

//Power is a control bool to be accessed to shut down the
//clientserver
var Power bool = true
var connectionsPerServer map[string]int = make(map[string]int)
var shutdownchan chan string
var logfile *os.File

func main() {
	var err error
	logfile, err = os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println(err)
	}

	go StartLoadBalancer("8082")

	go GrabServers()
	<-shutdownchan
}

//StartLoadBalancer begins the hosting process for the
//load balancer which assumes all incomming traffic is
//for the same type of server and routes messages to the
//least used server
func StartLoadBalancer(port string) {
	fmt.Println("Launching Load Balancing server...")
	logfile.Write([]byte("Launching Load Balancing server...\n"))

	ln, _ := net.Listen("tcp", ":"+port)

	fmt.Println("Online - Now Listening On Port: " + port)
	logfile.Write([]byte("Online - Now Listening On Port: " + port + "\n"))
	fmt.Println()

	ConnSignal := make(chan string)

	for Power {

		go Session(ln, ConnSignal, port)
		logfile.Write([]byte(<-ConnSignal))

	}
	fmt.Println("Shut Down Signal Sent...Ending")
	logfile.Write([]byte("Shut Down Signal Sent...Ending" + "\n"))
}

var shutDownSession chan string = make(chan string)

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

	//Matches conn to initial Server based on server conns
	//through the loadbalancer.
	var serverConn net.Conn = nil
	var intslice []int
	var err error

	for k := range connectionsPerServer {
		intslice = append(intslice, connectionsPerServer[k])
	}
	sort.Ints(intslice)
	valueToLookfor := intslice[0]
	for k, v := range connectionsPerServer {
		if valueToLookfor == v {
			serverConn, err = net.Dial("tcp", k)
			defer serverConn.Close()
			serverConn.Write(buf)
			logfile.Write([]byte("Connection Between I/O made\n"))
			logfile.Write([]byte("Message Sent: " + string(buf) + "\n"))
			if err != nil {
			}
			connectionsPerServer[k]++
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

			connectionsPerServer[conn] = 0
			fmt.Println("Added")

		case "distribute":
			ForceDistribution()
		}

	}

}

//ForceDistribution will examine connections and will
//redistribute connections to servers incase of clients
//mass disconnecting from one server. Only works with
//shared database for server connections, otherwise data
//would not transfer.
func ForceDistribution() {

}
