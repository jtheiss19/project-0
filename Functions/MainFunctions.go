//Package mainfunctions supplies functions for main.go.
package mainfunctions

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
	DB.SaveDB(DB.File)
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
	DB.SaveDB(DB.File)
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
	DB.SaveDB(DB.File)
}

//EndCol removes a new column to the database and then
//updates all other rows to remove that column's data
func EndCol(DelCol string, DB *DB.Database) {
	DB.DelCol(DelCol)
	DB.SaveDB(DB.File)
}

//CheckColHeader returns a bool based on if the database
//contains the column. NOT IN UNITTEST
func CheckColHeader(DB *DB.Database, ColHeader ...string) bool {
	HeaderSlice := DB.GetHeaders()

	Count := 0

	for i := 0; i < len(HeaderSlice); i++ {
		for j := 0; j < len(ColHeader); j++ {
			if HeaderSlice[i] == ColHeader[j] {
				Count++
			}
		}
	}
	if Count == len(ColHeader) {
		return true
	}
	return false
}

//CalculateBMI itterates through each row of a database and
//calculates the BMI for each row participant if possible. If
//there is a nil in a required field, returns nil into row
//NOT IN UNITTEST
func CalculateBMI(DB *DB.Database) {
	for i := 1; i < len(DB.Data); i++ {
		Weight, WError := strconv.ParseFloat(DB.Data[i][DB.GetColKey("Weight")], 32)
		Height, HError := strconv.ParseFloat(DB.Data[i][DB.GetColKey("Height")], 32)
		if WError == nil && HError == nil {
			BMIF := 703 * Weight / (Height * Height)
			BMI := fmt.Sprintf("%3.1f", BMIF)
			OverWriteCol(DB.Data[i][0], "BMI", BMI, DB)
		} else {
			fmt.Println("Check Database for none parsable stings in float lines: Height and Weight")
			return
		}
	}
	DB.SaveDB(DB.File)
}
