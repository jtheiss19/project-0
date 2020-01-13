//Package ezdb Provides simple easy database functionality to projects
package ezdb

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Database is struct for holding the database and cleaning up
//code functions that alter the data in the database
type Database struct {
	Data map[int][]string
	File string
}

//ReadDB is a function that reads Profiles.txt in the local
//folder and pulls all information stored in it into a mass
//slice of bytes. These byte slices can be used as a sortable
//field. Returns a new database.
func ReadDB(FileName string) *Database {

	// read in the contents of the localfile.data
	data, err := os.Open(FileName)

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

	Database := Database{Data: DB, File: FileName}

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
func (DB *Database) SaveDB(FileName string) {

	File, Error := os.OpenFile(FileName, os.O_TRUNC|os.O_WRONLY, 7777)
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

//DelRow removes a line at RowID in the Database's Data.
//Autoshifts the data to fill the spot for a continous
//indexing system.
func (DB *Database) DelRow(RowID int) {
	if RowID > len(DB.Data)-1 || RowID == 0 {
		return
	}

	for i := RowID; i+1 < len(DB.Data); i++ {
		DB.Data[i] = DB.Data[i+1]
		DB.Data[i][0] = (strconv.Itoa(i))
	}

	delete(DB.Data, len(DB.Data)-1)

}

//CreateCol adds a new Column to the header. Add
//"nil" into all other rows section under this
//new column
func (DB *Database) CreateCol(NewCol string) {
	DB.Data[0] = append(DB.Data[0], NewCol)

	for j := 1; j < len(DB.Data); j++ {
		DB.Data[j] = append(DB.Data[j], "nil")
	}
}

//DelCol removes a Column from the header. It also
//loops through all other rows to remove the value
//associated with that column
func (DB *Database) DelCol(DelCol string) {

	var Col int = 0
	for i := 0; i < len(DB.Data[0]); i++ {
		if DB.Data[0][i] == DelCol {
			Col = i
			break
		}
	}

	//Loops through row j
	for j := 0; j < len(DB.Data); j++ {
		//Loops through entries starting at Col: the entry being removed
		for i := Col; i < len(DB.Data[j]); i++ {
			if i+1 >= len(DB.Data[j]) {
				break
			}
			DB.Data[j][i] = DB.Data[j][i+1]
		}
		DB.Data[j] = DB.Data[j][0 : len(DB.Data[j])-1]
	}

}

//PrettyPrint allows the user to print the
//Data in the database object in a visually
//pleasing way into the console.
func (DB *Database) PrettyPrint() {

	for i := 0; i < len(DB.Data); i++ {

		for j := 0; j < len(DB.Data[i]); j++ {
			if j == 0 || i == 0 {
				fmt.Print(DB.Data[i][j])
			}

			for k := 0; k < 10-len(DB.Data[i][j]); k++ {
				fmt.Print(" ")
			}

			if i != 0 && j != 0 {
				fmt.Print(DB.Data[i][j])
			}

			fmt.Print(" | ")
			if j == len(DB.Data[i])-1 {
				fmt.Print("\n")
			}
		}
	}
}
