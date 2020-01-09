//Package db Provides database functionality to projects
package db

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Database is struct for holding the database and cleaning up
//code functions that alter the data in the database
type Database struct {
	Data map[int][]string
}

//ReadDB is a function that reads Profiles.txt in the local
//folder and pulls all information stored in it into a mass
//slice of bytes. These byte slices can be used as a sortable
//field. Returns a new database.
func ReadDB() *Database {

	// read in the contents of the localfile.data
	data, err := os.Open("Profiles.txt")

	//Error Handling
	if err != nil {
		fmt.Println(err)
	}

	reader := bufio.NewScanner(data)
	DB := make(map[int][]string)

	for i := 0; i >= 0; i++ {
		if reader.Scan() == false {
			break
		}
		Line := strings.Split(reader.Text(), ",")
		DB[i] = Line
	}

	Database := Database{Data: DB}

	//Try to close file
	Error := data.Close()
	//Error Handling
	if Error != nil {
		fmt.Println(Error)
		return &Database
	}

	return &Database
}

//SaveDB provides a method in which database's data can be
//saved, with alterations, into plain txt document for
//storage and editing
func (DB *Database) SaveDB() {

	File, Error := os.OpenFile("Profiles.txt", os.O_TRUNC|os.O_WRONLY, 7777)
	if Error != nil {
		fmt.Println(Error)
	}

	w := bufio.NewWriter(File)

	for i := 0; i < len(DB.Data); i++ {
		for j := 0; j < len(DB.Data[i]); j++ {
			w.WriteString(DB.Data[i][j])
			if j+1 == len(DB.Data[i]) {
				break
			}
			w.WriteString(",")
		}
		if i+1 == len(DB.Data) {
			break
		}
		w.WriteString("\n")
	}

	w.Flush()
	return
}

//GrabDBCol will search a Database and remove all columns
//except the key column and the column matching the ColTerm
//arg and output it as a new database
func (DB *Database) GrabDBCol(ColTerm string) *Database {

	var Col int = 0
	for i := 0; i < len(DB.Data[0]); i++ {
		if DB.Data[0][i] == ColTerm {
			Col = i
			break
		}
	}

	NewDB := make(map[int][]string)
	var AppendingArray []string
	for i := 0; i < len(DB.Data); i++ {
		AppendingArray = append(AppendingArray, DB.Data[i][Col])
		fmt.Println(AppendingArray)
		NewDB[i] = AppendingArray
		AppendingArray = nil
	}
	NewDatabase := Database{Data: NewDB}
	return &NewDatabase

}

//GrabDBRow returns a database which contains the headers
//row and the entirety of the row selected which matches
//the Key in RowID
func (DB *Database) GrabDBRow(RowID int) *Database {

	RowData := make(map[int][]string)

	RowData[0] = DB.Data[0]
	RowData[1] = DB.Data[int(RowID)]

	NewDatabase := Database{Data: RowData}
	return &NewDatabase
}

//GetRowKey retrieves the key when given the searchterm
//if found in any column
func (DB *Database) GetRowKey(SearchTerm string) int {

	var RowID int

	for i := 0; i < len(DB.Data); i++ {
		for j := 0; j < len(DB.Data[i]); j++ {
			if DB.Data[i][j] == SearchTerm {
				RowID = i
			}
		}
	}
	return RowID
}

//AddRow attaches the Line []string to the end of the
//Database's Data
func (DB *Database) AddRow(Line []string) {
	DB.Data[len(DB.Data)] = Line
}

//DelRow removes a line at RowID in the Database's Data
func (DB *Database) DelRow(RowID int) {
	delete(DB.Data, RowID)
}
