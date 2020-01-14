package main

import (
	"flag"
	"fmt"
	"os"

	Functions "github.com/jtheiss19/project-0/Functions"
)

func main() {

	Showcmd := flag.NewFlagSet("Show", flag.ExitOnError)
	ShowSpecify := Showcmd.Bool("s", false, "Toggles wheather a specific row will be showed. Must provide a search term.")

	Addcmd := flag.NewFlagSet("Add", flag.ExitOnError)
	AddCol := Addcmd.String("c", "", "Adds a new column to the current selected database")

	Delcmd := flag.NewFlagSet("Del", flag.ExitOnError)
	DelCol := Delcmd.String("c", "", "Removes a column from the current selected database")

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
			fmt.Println(Database.GrabDBRow(Key).PrettyPrint())
		} else {
			fmt.Println(Database.PrettyPrint())
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

	case "New":

	case "Review":

	case "Advice":
	}
}
