//Package database Provides database functionality to projects
package database

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Database is struct for holding the database and cleaning up code functions that alter the data in the database
type Database struct {
	Data [][]string
}

//ReadDB is a function that reads Profiles.txt in the local folder and pulls all information stored in it into a mass slice of bytes.
//These byte slices can be used as a sortable field. Returns a database of shape [][][]Byte
func ReadDB() *Database {

	// read in the contents of the localfile.data
	data, err := os.Open("Profiles.txt")

	//Error Handling
	if err != nil {
		fmt.Println(err)
	}

	reader := bufio.NewScanner(data)
	var DB [][]string

	for {
		if reader.Scan() == false {
			break
		}
		Line := strings.Split(reader.Text(), ",")
		DB = append(DB, Line)
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

//SaveDB provides a method to profile in which it's variables can be saved, with alterations, into plain txt document for storage and editing
func (DB *Database) SaveDB() {

	File, Error := os.OpenFile("access.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 7777)
	if Error != nil {
		fmt.Println(Error)
	}

	w := bufio.NewWriter(File)

	for i := 0; i < len(DB.Data); i++ {

		w.Write([]byte("test"))
		return
	}
	w.Write([]byte("test"))
	return
}

//GrabDBCol will search a DB of type [][][]byte and remove all columns except the key column and the column matching the ColTerm arg
func (DB *Database) GrabDBCol(ColTerm string) *Database {

	var Col int = 0
	for i := 0; i < len(DB.Data[0]); i++ {
		if DB.Data[0][i] == ColTerm {
			Col = i
			break
		}
	}

	if Col == 0 {
		fmt.Println("Could not find Column of name: " + ColTerm)
		return DB
	}

	var NewDB [][]string
	var AppendingArray []string
	for i := 0; i < len(DB.Data); i++ {
		AppendingArray = append(AppendingArray, DB.Data[i][0])
		AppendingArray = append(AppendingArray, DB.Data[i][Col])
		NewDB = append(NewDB, AppendingArray)
		AppendingArray = nil
	}
	NewDatabase := Database{Data: NewDB}
	return &NewDatabase

}

//GrabDBRow returns a database of type [][][]byte which contains the headers row (ID 0000) and the entirety of the row selected which matches the Key in RowID
func (DB *Database) GrabDBRow(RowID string) *Database {

	var RowData [][]string
	RowIDint, Error := strconv.ParseInt(RowID, 10, 32)
	if Error != nil {
		fmt.Println(Error)
		return DB
	}

	RowData = append(RowData, DB.Data[0])
	RowData = append(RowData, DB.Data[RowIDint])

	NewDatabase := Database{Data: RowData}
	return &NewDatabase
}

//GetRowKey retrieves the key when given the searchterm if found in any column
func (DB *Database) GetRowKey(SearchTerm string) string {

	var RowID string

	for i := 0; i < len(DB.Data); i++ {
		for j := 0; j < len(DB.Data[i]); j++ {
			if DB.Data[i][j] == SearchTerm {
				RowID = DB.Data[i][0]
			}
		}
	}
	return RowID
}
