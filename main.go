package main

import (
	"fmt"

	DB "github.com/jtheiss19/project-0/Database"
)

//AddProfile interfaces with the Database class to
//provide a special way to add data into it. This
//ensures formating is correct and matches the database
func AddProfile(DB DB.Database) {

}

//DelProfile interfaces with the Database class to
//provide a special way to remove data from it. This
//provides extra functionality in how a profile is removed
func DelProfile(DB DB.Database) {

}

func main() {
	Database := DB.ReadDB()
	Database.DelRow(4)
	fmt.Println(Database.Data)
	Database.SaveDB()

}
