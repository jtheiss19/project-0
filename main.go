package main

import (
	"flag"
	"fmt"
	"os"
)

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
