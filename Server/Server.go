//Package server is a test package not for final product
package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
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

var Power bool = true
var connections []net.Conn

//StartClientServer begins the hosting process for the
//client to server application
func StartClientServer(Database *EZDB.Database, port string) {
	fmt.Println("Launching server...")

	ln, _ := net.Listen("tcp", ":"+port)

	fmt.Println("Online - Now Listening On Port: " + port)
	ConnSignal := make(chan string)

	for Power {

		go Session(ln, Database, ConnSignal, port)
		<-ConnSignal

	}
	for i := 0; i < len(connections); i++ {
		fmt.Fprintf(connections[i], "END"+string('\u0000'))
		//fmt.Println(connections[i])
	}
}

func Session(ln net.Listener, Database *EZDB.Database, ConnSignal chan string, port string) {
	conn, _ := ln.Accept()
	connections = append(connections, conn)
	fmt.Println("New Connection On Port: " + port)
	ConnSignal <- "New Connection"
	for {

		message, _ := bufio.NewReader(conn).ReadString('\n')

		fmt.Print("Command Received: ", string(message))

		Command := strings.Split(string(message), " ")

		fmt.Fprintf(conn, string('\n'))

		switch Command[0] {

		case "Show":
			if strings.Contains(string(message), "-s") {
				Key := Database.GetRowKey(string(Command[2]))
				Information := (Database.GrabDBRow(Key).PrettyPrint())
				for i := 0; i < len(Information); i++ {
					fmt.Fprintf(conn, Information[i]+string('\n'))
				}
			} else {
				for i := 0; i < len(Database.PrettyPrint()); i++ {
					fmt.Fprintf(conn, Database.PrettyPrint()[i]+string('\n'))
				}

			}

		case "Disconnect":
			fmt.Println("Connection Terminated")
			fmt.Fprintf(conn, "Connection Terminated"+string('\n'))
			return

		case "Power":
			fmt.Println("Connection Terminated - Powering Down")
			Power = false
			ConnSignal <- "End"

		case "":
			fmt.Println("Connection Terminated")
			fmt.Fprintf(conn, "Connection Terminated"+string('\n'))
			return

		default:
			fmt.Println("Could not understand: " + string(message))
			fmt.Fprintf(conn, "Not Parseable"+string('\n'))
		}
		fmt.Fprintf(conn, string('\u0000'))

	}
}
