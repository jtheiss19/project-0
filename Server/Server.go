//Package server is a test package not for final product
package server

import (
	"fmt"
	"net/http"
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

//StartServer begins the hosting process for the
//webserver
func StartServer(DB *EZDB.Database) {

	Database = DB

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("Server/WebPages"))))
	http.HandleFunc("/", Handler)
	fmt.Println("Online - Now Listening")

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Println(err)
	}
}
