//Package server is a test package not for final product
package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"text/template"

	EZDB "github.com/jtheiss19/project-0/Database"
)

//Database global
var Database *EZDB.Database

//Handler is main handler for running the webpage.
//It passes the global variable Database for parsing
//index.html for pulling data.
func Handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("Server/WebPages/index.html")
	t.Execute(w, Database)
}

//PatientHandler is for handling a patients profile.
//It passes the global variable Database for parsing
//patient.html for pulling data.
func PatientHandler(w http.ResponseWriter, r *http.Request) {
	keys, _ := r.URL.Query()["key"]
	if len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}
	keyint := 0
	for i := 0; i < len(keys); i++ {
		key := keys[i]
		keyint2, _ := strconv.Atoi(key)
		keyint = keyint + keyint2
	}

	if keyint >= len(Database.Data) {
		keyint = len(Database.Data) - 1
	}
	if keyint <= 0 {
		keyint = 1
	}

	PatientData := Database.GrabDBRow(keyint)

	t, _ := template.ParseFiles("Server/WebPages/patient.html")
	t.Execute(w, PatientData)
}

//StartHTMLServer begins the hosting process for the
//webserver
func StartHTMLServer(DB *EZDB.Database, port string) {

	Database = DB

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("Server/WebPages"))))
	http.HandleFunc("/", Handler)
	http.HandleFunc("/view/", PatientHandler)
	fmt.Println("Online - Now Listening On Port: " + port)

	err := http.ListenAndServe("localhost:"+port, nil)
	if err != nil {
		fmt.Println(err)
	}
}

//Power is a control bool to be accessed to shut down the
//clientserver
var Power bool = true
var connections []net.Conn

//StartClientServer begins the hosting process for the
//client to server application
func StartClientServer(Database *EZDB.Database, port string) {
	fmt.Println("Launching server...")

	ln, _ := net.Listen("tcp", ":"+port)

	fmt.Println("Online - Now Listening On Port: " + port)
	fmt.Println()

	ConnSignal := make(chan string)

	for Power {

		go Session(ln, Database, ConnSignal, port)
		<-ConnSignal

	}
	fmt.Println("Shutting Down...")
	for i := 0; i < len(connections); i++ {
		fmt.Fprintf(connections[i], "Server is shutting down, Disconnecting you \n"+string('\u0007')+string('\u0000'))
	}
	fmt.Println("Shut Down Signal Sent...Ending")
}

//Session creates a new seesion listening on a port. This
//session handles all interactions with the connected
//client
func Session(ln net.Listener, Database *EZDB.Database, ConnSignal chan string, port string) {
	conn, _ := ln.Accept()
	connections = append(connections, conn)

	host, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	fmt.Println("New Connection On Port: " + port + " from " + host)
	fmt.Println()
	ConnSignal <- "New Connection On Port: " + port + " from " + host

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
			log.Fatal(err)
		}

		fmt.Fprintf(conn, "%s\n", out)

		fmt.Fprintf(conn, string('\u0000'))
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

		message, _ := bufio.NewReader(conn).ReadString('\n')

		command := strings.TrimRight(string(message), " \n")
		command = strings.TrimSpace(string(message))
		commandSlice := strings.Split(command, " ")

		fmt.Print("Raw Text Received From "+host+": ", string(message))

		if command == "Power" {
			Power = false
			ConnSignal <- "Remoted Power Toggled from "
			return
		}

		if command == "Disconnect" {
			fmt.Fprintf(conn, "Disconnecting you from Server \n"+string('\u0007')+string('\u0000'))
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

		fmt.Println("Command to exectute for "+host+": ", command)
		fmt.Println()

		messages <- commandSlice

	}
}
