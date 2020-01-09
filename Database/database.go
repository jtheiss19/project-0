//Package database Provides database functionality to projects
package database

import (
	"fmt"
	"io/ioutil"
	"math"
)

//ByteToFloat provides functionality to convert a slice of bytes into a float32 using unicode's number for zero(48) as an offset.
//The tens place is held by the offset var which tracks the position of the decimal point.
func ByteToFloat(Byte []byte) float32 {
	var tempval float32 = 0
	var offset int = 0
	for i := 0; i < len(Byte); i++ {
		if string(Byte[i]) == "." {
			offset = -1
		} else {
			tempval = tempval + (float32(Byte[i])-48)*float32(math.Pow10(len(Byte)-i-offset))
		}
	}
	return tempval / float32(math.Pow10(len(Byte)+offset))
}

//ReadDB is a function that reads Profiles.txt in the local folder and pulls all information stored in it into a mass slice of bytes.
//These byte slices can be used as a sortable field. Returns a database of shape [][][]Byte
func ReadDB() [][][]byte {

	// read in the contents of the localfile.data
	data, err := ioutil.ReadFile("Profiles.txt")

	//Error Handling
	if err != nil {
		fmt.Println(err)
	}

	//Var DB is the overall database
	//Var Line holds the line that is currently being pulled into so that it may be appended into DB
	//Var Size holds len(data)
	var DB [][][]byte
	var Line [][]byte
	var Size int = len(data)

	i, k := -1, 0
	//Loop through Lines
	for i+1 < Size && k < Size {
		//Loop through entries in Lines
		for i < Size && k < Size {
			//Alternates between two loops which capture the length of the entry by leap frogging l and k over eachother
			for k = i + 1; k < Size; k++ {
				if data[k] == '\u002C' || data[k] == '\u000a' {
					Line = append(Line, data[i+1:k])
					break
				}
			}
			for i = k + 1; i < Size; i++ {
				if data[i] == '\u002C' || data[i] == '\u000a' {
					Line = append(Line, data[k+1:i])
					break
				}
			}

			//Provies a break out of the loop that searches through the line for entries
			if data[i] == '\u000a' {
				break
			}
		}
		DB = append(DB, Line)
		Line = nil

	}
	return DB
}

//GrabDBCol will search a DB of type [][][]byte and remove all columns except the key column and the column matching the ColTerm arg
func GrabDBCol(ColTerm string, DB [][][]byte) [][][]byte {

	var Col int = 0
	for i := 0; i < len(DB[0]); i++ {
		if string(DB[0][i]) == ColTerm {
			Col = i
			break
		}
	}

	if Col == 0 {
		fmt.Println("Could not find Column of name: " + ColTerm)
		return DB
	}

	var NewDB [][][]byte
	var AppendingArray [][]byte
	for i := 0; i < len(DB); i++ {
		AppendingArray = append(AppendingArray, DB[i][0])
		AppendingArray = append(AppendingArray, DB[i][Col])
		NewDB = append(NewDB, AppendingArray)
		AppendingArray = nil
	}
	return NewDB

}

//GrabDBRow returns a database of type [][][]byte which contains the headers row (ID 0000) and the entirety of the row selected which matches the Key in RowID
func GrabDBRow(RowID []byte, DB [][][]byte) [][][]byte {
	IDint := int(ByteToFloat(RowID))
	var RowData [][][]byte
	RowData = append(RowData, DB[0])
	RowData = append(RowData, DB[IDint])
	return RowData
}

//
func GetRowKey(SearchTerm string, DB [][][]byte) []byte {

	for i := 0; i < len(DB[0]); i++ {
		if string(DB[0][i]) == ColTerm {
			Col = i
			break
		}
	}

	var test []byte
	return test
}

//DisplayDBAsString is a function that converts a database of type [][][]byte to a 2-D slice of strings that allow for easy viewing of the database
func DisplayDBAsString(DB [][][]byte) [][]string {
	var NewDB [][]string
	var NewLine []string
	var Entry string
	for i := 0; i < len(DB); i++ {
		for j := 0; j < len(DB[i]); j++ {
			Entry = string(DB[i][j])
			NewLine = append(NewLine, Entry)
		}
		NewDB = append(NewDB, NewLine)
		NewLine = nil

	}
	return NewDB
}
