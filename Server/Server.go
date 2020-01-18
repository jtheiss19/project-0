//Package server is a test package not for final product
package server

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	EZDB "github.com/jtheiss19/project-0/Database"
	Functions "github.com/jtheiss19/project-0/Functions"
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

//StartClientServer begins the hosting process for the
//client to server application
func StartClientServer(Database *EZDB.Database, port string) {
	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":"+port)

	fmt.Println("Online - Now Listening On Port: " + port)

	// accept connection on port
	conn, _ := ln.Accept()

	fmt.Println("New Connection On Port: " + port)

	Showcmd := flag.NewFlagSet("Show", flag.ExitOnError)
	ShowSpecify := Showcmd.Bool("s", false, "Toggles wheather a specific row will be showed. Must provide a search term.")

	Addcmd := flag.NewFlagSet("Add", flag.ExitOnError)
	AddCol := Addcmd.String("c", "", "Adds a new column to the current selected database")

	Delcmd := flag.NewFlagSet("Del", flag.ExitOnError)
	DelCol := Delcmd.String("c", "", "Removes a column from the current selected database")

	Reviewcmd := flag.NewFlagSet("Review", flag.ExitOnError)
	ReviewPerson := Reviewcmd.String("p", "", "Looks up the profile by key and reviews health information if available.")
	// run loop forever (or until ctrl-c)
	for {
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		fmt.Print("Command Received: ", string(message))

		Command := strings.Split(string(message), " ")

		switch Command[0] {

		case "Test":
			fmt.Fprintf(conn, string(message))

		case "Add":
			Addcmd.Parse(Command[1:])
			if *AddCol != "" {
				Functions.NewCol(*AddCol, Database)
			} else {
				Functions.AddProfile(Command[1:], Database)
			}

		case "Del":
			Delcmd.Parse(Command[1:])
			if *DelCol != "" {
				Functions.EndCol(*DelCol, Database)
			} else {
				Functions.DelProfile(Command[1], Database)
			}

		case "Edit":
			Functions.OverWriteCol(Command[1], Command[2], Command[3], Database)

		case "Replace":
			Functions.Replace(Command[1], Command[2:], Database)

		case "Show":
			Showcmd.Parse(Command)
			if *ShowSpecify {
				Key := Database.GetRowKey(Command[2])
				Information := (Database.GrabDBRow(Key).PrettyPrint())
				for i := 0; i < len(Information); i++ {
					fmt.Println(Information[i])
				}
			} else {
				for i := 0; i < len(Database.PrettyPrint()); i++ {
					fmt.Println(Database.PrettyPrint()[i])
				}

			}

		case "Calc":
			if Functions.CheckColHeader(Database, "Weight", "Height") {
				if !Functions.CheckColHeader(Database, "BMI") {
					Database.CreateCol("BMI")
				}
				Functions.CalculateBMI(Database)
			} else {
				fmt.Println("Missing Columns to calculate BMI")
			}

		case "Switch":
			Database = EZDB.ReadDB("Database/Databases/" + string(Command[1]) + ".txt")

		case "Review":
			Reviewcmd.Parse(Command[1:])
			if *ReviewPerson != "" {
				Functions.Review(*ReviewPerson, Database, "Database/Databases/BMI.txt")
			} else {
				fmt.Println("Need to specify profile")
			}

		case "Exit":
			fmt.Println("Connection Terminated")
			break

		default:
			fmt.Println("Could not understand: " + string(message))
			fmt.Fprintf(conn, "Not Parseable"+"\n")
		}

	}
}
