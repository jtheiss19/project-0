package main

import (
	"fmt"
	"os"

	database "github.com/jtheiss19/project-0/Database"
)

//Profile acts as an object which can hold the various internal variables of a profile while loaded when operations are applied.
type Profile struct {
	Name string
	Age  int
}

//Save provides a method to profile in which it's variables can be saved, with alterations, into plain txt document for storage and editing
func (P Profile) Save() {
	//Try to open File with Profiles.txt
	File, Error := os.OpenFile("Profiles.txt", os.O_APPEND|os.O_WRONLY, 0644)
	//Error Handling
	if Error != nil {
		fmt.Println(Error)
		return
	}

	//Try to wrtie into File
	_, Error = fmt.Fprintln(File, P.Name)
	//Error Handling
	if Error != nil {
		fmt.Println(Error)
		return
	}

	//Try to close file
	Error = File.Close()
	//Error Handling
	if Error != nil {
		fmt.Println(Error)
		return
	}

	return
}

func main() {
	Database := database.ReadDB()
	Database = database.GrabDBCol("Names", Database)
	fmt.Println(database.DisplayDBAsString(Database))

}
