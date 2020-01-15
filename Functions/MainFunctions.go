//Package mainfunctions supplies functions for main.go.
package mainfunctions

import (
	"fmt"
	"strconv"

	EZDB "github.com/jtheiss19/project-0/Database"
)

//AddProfile interfaces with the Database class to
//provide a special way to add data into it. This
//ensures formating is correct and matches the database
func AddProfile(UserInput []string, DB *EZDB.Database) {
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
func DelProfile(ProfileID string, DB *EZDB.Database) {
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
func OverWriteCol(ProfileID string, Column string, NewVal string, DB *EZDB.Database) {

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
func Replace(ID string, NewProfile []string, DB *EZDB.Database) {
	DelProfile(ID, DB)
	AddProfile(NewProfile, DB)

}

//NewCol adds a new column to the database and then
//updates all other rows to include none in the column
func NewCol(NewCol string, DB *EZDB.Database) {
	DB.CreateCol(NewCol)
	DB.SaveDB(DB.File)
}

//EndCol removes a new column to the database and then
//updates all other rows to remove that column's data
func EndCol(DelCol string, DB *EZDB.Database) {
	DB.DelCol(DelCol)
	DB.SaveDB(DB.File)
}

//CheckColHeader returns a bool based on if the database
//contains the column. NOT IN UNITTEST
func CheckColHeader(DB *EZDB.Database, ColHeader ...string) bool {
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
func CalculateBMI(DB *EZDB.Database) {
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

//Review looks up a profile in the database given is ID
//and review the health information to tell someone how
//healthy they are and provided solutions for health
//improvement. NOT IN UNITTEST
func Review(ProfileID string, DB *EZDB.Database, BMITable string) {
	var Catagory int
	Key, _ := strconv.Atoi(ProfileID)

	NewDB := DB.GrabDBRow(Key)
	NewDB.DelCol("Key")
	fmt.Println(NewDB.PrettyPrint())

	BMIData := EZDB.ReadDB(BMITable)
	BMIData.DelCol("Key")

	HKey := DB.GetColKey("Height")
	HeightF, _ := strconv.ParseFloat(DB.Data[Key][HKey], 32)

	BMIKey := DB.GetColKey("BMI")
	BMII, _ := strconv.ParseFloat(DB.Data[Key][BMIKey], 32)

	BMIMinKey := BMIData.GetColKey("BMI Min")
	BMIMaxKey := BMIData.GetColKey("BMI Max")

	for i := 1; i < len(BMIData.Data); i++ {
		BMIMin, _ := strconv.ParseFloat(BMIData.Data[i][BMIMinKey], 32)
		BMIMax, _ := strconv.ParseFloat(BMIData.Data[i][BMIMaxKey], 32)

		CalculatedWeightMin := HeightF * HeightF * BMIMin / 703
		CalculatedWeightMax := HeightF * HeightF * BMIMax / 703

		CalculatedWeightMinS := fmt.Sprintf("%.1f", CalculatedWeightMin)
		CalculatedWeightMaxS := fmt.Sprintf("%.1f", CalculatedWeightMax)

		OverWriteCol(strconv.Itoa(i), "Weight Min", CalculatedWeightMinS, BMIData)
		OverWriteCol(strconv.Itoa(i), "Weight Max", CalculatedWeightMaxS, BMIData)

		fmt.Print("|")
		for j := int(BMIMin); j < int(BMIMax-BMIMin)/2+int(BMIMin); j++ {
			if int(BMII) == j {
				fmt.Print("X")
				Catagory = i
			}
			fmt.Print("-")
		}
		fmt.Print(BMIData.Data[i][0])
		for j := int(BMIMax-BMIMin)/2 + int(BMIMin); j < int(BMIMax); j++ {
			if int(BMII) == j {
				fmt.Print("X")
				Catagory = i
			}
			fmt.Print("-")
		}
	}
	fmt.Print("| \n \n")
	fmt.Print("They are currently considered ", BMIData.Data[Catagory][0], ". ")
	if Catagory > 2 {
		fmt.Println("Consider a diet to bring their weight back into the normal range.")
	} else if Catagory < 2 {
		fmt.Println("Have them consider eating more. Being underweight is not healthy.")
	} else {
		fmt.Println("Keep up the good work. Do not let them waste this good health!")
	}
	fmt.Println()
	fmt.Println(BMIData.PrettyPrint())
}
