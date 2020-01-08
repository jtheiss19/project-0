package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
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

//Profile acts as an object which can hold the various internal variables of a profile while loaded when operations are applied.
type Profile struct {
	Name string
	Age  int
}

//Save provides a method to profile in which it's variables can be saved, with alterations, into plain txt document for storage and editing
func (P Profile) Save() {
	//Try to open File with Profiles.txt
	File, Error := os.OpenFile("Profiles.txt", os.O_APPEND|os.O_WRONLY, 0644)
	//Error Handling
	if Error != nil {
		fmt.Println(Error)
		return
	}

	//Try to wrtie into File
	_, Error = fmt.Fprintln(File, P.Name)
	//Error Handling
	if Error != nil {
		fmt.Println(Error)
		return
	}

	//Try to close file
	Error = File.Close()
	//Error Handling
	if Error != nil {
		fmt.Println(Error)
		return
	}

	return
}

//ReadProfiles is a function that reads Profiles.txt in the local folder and pulls all information stored in it into a mass slice of bytes.
//These byte slices can be used as a sortable field. Returns a database of shape [][][]Byte
func ReadProfiles() [][][]byte {

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

	i, k := 0, 0
	//Loop through Lines
	for i+1 < Size && k < Size {
		//Loop through entries in Lines
		for i < Size && k < Size {
			//Alternates between two loops which capture the length of the entry by leap frogging l and k over eachother
			for k = i + 1; k < Size; k++ {
				if string(data[k]) == "," || string(data[k]) == "}" {
					Line = append(Line, data[i+1:k])
					break
				}
			}
			for i = k + 1; i < Size; i++ {
				if string(data[i]) == "," || string(data[i]) == "}" {
					Line = append(Line, data[k+1:i])
					break
				}
			}

			//Provies a break out of the loop that searches through the line for entries
			if string(data[i]) == "}" {
				break
			}
		}
		DB = append(DB, Line)
		Line = nil
	}
	fmt.Println(DB)
	return DB
}

func main() {
	ReadProfiles()
}
