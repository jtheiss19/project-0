package main

import (
	"fmt"
	"strconv"

	DB "github.com/jtheiss19/project-0/Database"
)

//AddProfile interfaces with the Database class to
//provide a special way to add data into it. This
//ensures formating is correct and matches the database
func AddProfile(UserInput []string, DB *DB.Database) {
	if len(DB.Data[0]) != len(UserInput) {
		fmt.Println("Not enough args have been entered")
		return
	}
	DB.AddRow(UserInput)
	Database.SaveDB(Config.Database)
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
	Database.SaveDB(Config.Database)
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

//NewCol adds a new column to the database and then
//updates all other rows to include none in the column
func NewCol(NewCol string, DB *DB.Database) {
	DB.CreateCol(NewCol)
	DB.SaveDB(Config.Database)
}

//EndCol adds a new column to the database and then
//updates all other rows to include none in the column
func EndCol(DelCol string, DB *DB.Database) {
	DB.DelCol(DelCol)
	DB.SaveDB(Config.Database)
}
