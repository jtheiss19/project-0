package main

import (
	"encoding/json"
	"fmt"
	"os"

	EZDB "github.com/jtheiss19/project-0/Database"
)

//Configuration struct holds the config settings
//pulled from the config file
type Configuration struct {
	Database string `json:"database"`
	BMITable string `json:"BMITable"`
	Port     string `json:"Port"`
}

//Database is the working database that this program
//will deal with. It is the main, unaltered database
//that should be saved when the program exits
var Database *EZDB.Database

//Config is the one and only iteration of the Configuration
//struct. It alone holds the sessions configuration settings
//from the config file which is loaded in during init()
var Config = Configuration{}

func init() {

	// read in the contents of the localfile.data
	File, Error := os.Open("config.json")
	//Error Handling
	if Error != nil {
		fmt.Println(Error)
		fmt.Println("Could not find file")
	}

	//Read config file
	Error = json.NewDecoder(File).Decode(&Config)
	//Error Handling
	if Error != nil {
		fmt.Println(Error)
		fmt.Println("Could not find file")
	}

	Database = EZDB.ReadDB(Config.Database)
}
