package main

import (
	"fmt"
	"os"
	"strconv"

	DB "github.com/jtheiss19/project-0/Database"
)

//Database is the working database that this program
//will deal with. It is the main, unaltered database
//that should be saved when the program exits
var Database *DB.Database

func init() {
	Database = DB.ReadDB()
}

//AddProfile interfaces with the Database class to
//provide a special way to add data into it. This
//ensures formating is correct and matches the database
func AddProfile(UserInput []string, DB *DB.Database) {
	if len(DB.Data[0]) != len(UserInput) {
		fmt.Println("Not enough args have been entered")
		return
	}
	DB.AddRow(UserInput)
	Database.SaveDB()
}

//DelProfile interfaces with the Database class to
//provide a special way to remove data from it. This can
//provide extra functionality in how a profile is removed
func DelProfile(ProfileID string, DB *DB.Database) {
	ID, Error := strconv.Atoi(ProfileID)
	if Error != nil {
		fmt.Println(Error)
		return
	}
	DB.DelRow(ID)
	Database.SaveDB()
}

func main() {
	switch os.Args[1] {

	case "Add":
		AddProfile(os.Args[2:], Database)

	case "Del":
		DelProfile(os.Args[2], Database)
	}
}
