package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	Functions "github.com/jtheiss19/project-0/Functions"
	Server "github.com/jtheiss19/project-0/Server"
)

func main() {

	Showcmd := flag.NewFlagSet("Show", flag.ExitOnError)
	ShowSpecify := Showcmd.Bool("s", false, "Toggles wheather a specific row will be showed. Must provide a search term.")

	Addcmd := flag.NewFlagSet("Add", flag.ExitOnError)
	AddCol := Addcmd.String("c", "", "Adds a new column to the current selected database")

	Delcmd := flag.NewFlagSet("Del", flag.ExitOnError)
	DelCol := Delcmd.String("c", "", "Removes a column from the current selected database")

	Reviewcmd := flag.NewFlagSet("Review", flag.ExitOnError)
	ReviewPerson := Reviewcmd.String("p", "", "Looks up the profile by key and reviews health information if available.")

	Hostcmd := flag.NewFlagSet("Host", flag.ExitOnError)
	HostHTML := Hostcmd.Bool("html", false, "Host an html server instead of a client server")

	if len(os.Args) < 2 {
		return
	}

	switch os.Args[1] {

	case "Add":
		Addcmd.Parse(os.Args[2:])
		if *AddCol != "" {
			Functions.NewCol(*AddCol, Database)
		} else {
			Functions.AddProfile(os.Args[2:], Database)
		}

	case "Del":
		Delcmd.Parse(os.Args[2:])
		if *DelCol != "" {
			Functions.EndCol(*DelCol, Database)
		} else {
			Functions.DelProfile(os.Args[2], Database)
		}

	case "Edit":
		Functions.OverWriteCol(os.Args[2], os.Args[3], os.Args[4], Database)

	case "Replace":
		Functions.Replace(os.Args[2], os.Args[3:], Database)

	case "Show":
		Showcmd.Parse(os.Args[2:])
		if *ShowSpecify {
			Key := Database.GetRowKey(os.Args[3])
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
		Config.Database = "Database/Databases/" + string(os.Args[2]) + ".txt"
		SaveFile, _ := json.MarshalIndent(Config, "", "	")
		_ = ioutil.WriteFile("config.json", SaveFile, 0644)

	case "Review":
		Reviewcmd.Parse(os.Args[2:])
		if *ReviewPerson != "" {
			Functions.Review(*ReviewPerson, Database, Config.BMITable)
		} else {
			fmt.Println("Need to specify profile")
		}

	case "Host":
		Hostcmd.Parse(os.Args[2:])
		if *HostHTML {
			Server.StartHTMLServer(Database, Config.Port)
		} else {
			if len(os.Args) > 2 {
				Server.StartClientServer(os.Args[2])
			} else {
				Server.StartClientServer(Config.Port)
			}

		}

	}

}
