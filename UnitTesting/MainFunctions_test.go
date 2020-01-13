package tests

import (
	"testing"

	EZDB "github.com/jtheiss19/project-0/Database"
	Functions "github.com/jtheiss19/project-0/Functions"
)

func TestNewCol(test *testing.T) {

	Data := make(map[int][]string)

	Data[0] = append(Data[0], "Key")
	Data[1] = append(Data[1], "0001")

	DB := &EZDB.Database{Data: Data, File: "SaveFile.txt"}

	Functions.NewCol("testing", DB)

	if DB.Data[0][len(DB.Data[0])-1] != "testing" {
		test.Errorf("New Column was expected to be testing. Intead got " + DB.Data[0][len(DB.Data[0])-1])
	}
	if DB.Data[1][len(DB.Data[0])-1] != "nil" {
		test.Errorf("New Column was expected to be nil. Intead got " + DB.Data[0][len(DB.Data[0])-1])
	}
}

func TestEndCol(test *testing.T) {

	Data := make(map[int][]string)

	Data[0] = append(Data[0], "Key")
	Data[0] = append(Data[0], "testing1")
	Data[0] = append(Data[0], "testing2")
	Data[1] = append(Data[1], "0001")
	Data[1] = append(Data[1], "testdata1")
	Data[1] = append(Data[1], "testdata2")

	DB := &EZDB.Database{Data: Data, File: "SaveFile.txt"}

	Functions.EndCol("testing2", DB)

	if DB.Data[0][len(DB.Data[0])-1] != "testing1" {
		test.Errorf("New Column was expected to be testing1. Intead got " + DB.Data[0][len(DB.Data[0])-1])
	}
	if DB.Data[1][len(DB.Data[0])-1] != "testdata1" {
		test.Errorf("New Column was expected to be testdata1. Intead got " + DB.Data[0][len(DB.Data[0])-1])
	}
}
