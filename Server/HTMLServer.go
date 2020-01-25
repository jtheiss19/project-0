//Package server is a test package not for final product
package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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
