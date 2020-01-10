package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	DB "github.com/jtheiss19/project-0/Database"
)

//Database is the working database that this program
//will deal with. It is the main, unaltered database
//that should be saved when the program exits
var Database *DB.Database
var ConfigFile = make(map[string]string)

func init() {
	// read in the contents of the localfile.data
	File, Error := os.Open("CONFIG.txt")

	//Error Handling
	if Error != nil {
		fmt.Println(Error)
		fmt.Println("Could not find file")
	}

	//Read config file
	reader := bufio.NewScanner(File)

	for i := 0; i >= 0; i++ {
		if reader.Scan() == false {
			break
		}
		Line := strings.Split(reader.Text(), "=")
		ConfigFile[Line[0]] = Line[1]
	}
	Database = DB.ReadDB(ConfigFile["DATABASE"])
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
	Database.SaveDB(ConfigFile["DATABASE"])
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
	Database.SaveDB(ConfigFile["DATABASE"])
}

//OverWriteCol is one method for replacing a single
//data entry in a row.
func OverWriteCol(ProfileID string, Column string, NewVal string, DB *DB.Database) {

	var Col int = 0
	for i := 0; i < len(DB.Data[0]); i++ {
		if DB.Data[0][i] == Column {
			Col = i
			break
		}
	}

	ID, Error := strconv.Atoi(ProfileID)
	if Error != nil {
		fmt.Println(Error)
		return
	}

	DB.Data[ID][Col] = NewVal

}

//Replace is one method for replacing data in a row.
//It works by rebuilding the entire profile.
func Replace(ID string, NewProfile []string, DB *DB.Database) {
	DelProfile(ID, DB)
	AddProfile(NewProfile, DB)

}

func main() {
	Showcmd := flag.NewFlagSet("Show", flag.ExitOnError)
	ShowSpecify := Showcmd.Bool("s", false, "Toggles wheather a specific row will be showed. Must provide a search term.")

	if len(os.Args) < 2 {
		return
	}

	switch os.Args[1] {

	case "Add":
		AddProfile(os.Args[2:], Database)

	case "Del":
		DelProfile(os.Args[2], Database)

	case "Edit":
		OverWriteCol(os.Args[2], os.Args[3], os.Args[4], Database)

	case "Replace":
		Replace(os.Args[2], os.Args[3:], Database)

	case "Show":
		Showcmd.Parse(os.Args[2:])
		if *ShowSpecify {
			Key := Database.GetRowKey(os.Args[3])
			fmt.Println(Database.GrabDBRow(Key))
		} else {
			fmt.Println(Database.Data)
		}

	}
}
