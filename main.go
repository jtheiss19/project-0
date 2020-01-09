package main

import (
	database "github.com/jtheiss19/project-0/Database"
)

//Profile acts as an object which can hold the various internal variables of a profile while loaded when operations are applied.
type Profile struct {
	Name string
	Age  int
}

func main() {
	Database := database.ReadDB()
	Database.SaveDB()
}
