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

func TestAddProfile(test *testing.T) {

	Data := make(map[int][]string)

	Data[0] = append(Data[0], "Key")
	Data[0] = append(Data[0], "testing1")
	Data[0] = append(Data[0], "testing2")
	Data[1] = append(Data[1], "0001")
	Data[1] = append(Data[1], "testdata1")
	Data[1] = append(Data[1], "testdata2")

	DB := &EZDB.Database{Data: Data, File: "SaveFile.txt"}

	var NewTestSlice []string

	NewTestSlice = append(NewTestSlice, "0002")
	NewTestSlice = append(NewTestSlice, "testcol1")
	NewTestSlice = append(NewTestSlice, "testcol2")

	Functions.AddProfile(NewTestSlice, DB)

	if DB.Data[2][0] != "0002" {
		test.Errorf("New Column was expected to be 0002. Intead got " + DB.Data[2][0])
	}
	if DB.Data[2][1] != "testcol1" {
		test.Errorf("New Column was expected to be testcol1. Intead got " + DB.Data[2][1])
	}
	if DB.Data[2][2] != "testcol2" {
		test.Errorf("New Column was expected to be testcol2. Intead got " + DB.Data[2][2])
	}

}

func TestDelProfile(test *testing.T) {
	Data := make(map[int][]string)

	Data[0] = append(Data[0], "Key")
	Data[0] = append(Data[0], "testing1")
	Data[0] = append(Data[0], "testing2")
	Data[1] = append(Data[1], "0001")
	Data[1] = append(Data[1], "testdata1")
	Data[1] = append(Data[1], "testdata2")
	Data[2] = append(Data[2], "0002")
	Data[2] = append(Data[2], "testcol1")
	Data[2] = append(Data[2], "testcol2")

	DB := &EZDB.Database{Data: Data, File: "SaveFile.txt"}

	Functions.DelProfile("1", DB)

	if DB.Data[0][0] != "Key" {
		test.Errorf("New Column was expected to be Key. Intead got " + DB.Data[0][0])
	}
	if DB.Data[0][1] != "testing1" {
		test.Errorf("New Column was expected to be testing1. Intead got " + DB.Data[0][1])
	}
	if DB.Data[0][2] != "testing2" {
		test.Errorf("New Column was expected to be testing2. Intead got " + DB.Data[0][2])
	}
	if DB.Data[1][0] != "1" {
		test.Errorf("New Column was expected to be 1. Intead got " + DB.Data[1][0])
	}
	if DB.Data[1][1] != "testcol1" {
		test.Errorf("New Column was expected to be testcol1. Intead got " + DB.Data[1][1])
	}
	if DB.Data[1][2] != "testcol2" {
		test.Errorf("New Column was expected to be testcol2. Intead got " + DB.Data[1][2])
	}
}

func TestOverWriteCol(test *testing.T) {
	Data := make(map[int][]string)

	Data[0] = append(Data[0], "Key")
	Data[0] = append(Data[0], "testing1")
	Data[0] = append(Data[0], "testing2")
	Data[1] = append(Data[1], "0001")
	Data[1] = append(Data[1], "testdata1")
	Data[1] = append(Data[1], "testdata2")
	Data[2] = append(Data[2], "0002")
	Data[2] = append(Data[2], "testcol1")
	Data[2] = append(Data[2], "testcol2")

	DB := &EZDB.Database{Data: Data, File: "SaveFile.txt"}

	Functions.OverWriteCol("1", "testing1", "NewData", DB)

	if DB.Data[1][1] != "NewData" {
		test.Errorf("New Column was expected to be NewData. Intead got " + DB.Data[1][0])
	}

}

func TestReplace(test *testing.T) {
	Data := make(map[int][]string)

	Data[0] = append(Data[0], "Key")
	Data[0] = append(Data[0], "testing1")
	Data[0] = append(Data[0], "testing2")
	Data[1] = append(Data[1], "1")
	Data[1] = append(Data[1], "testdata1")
	Data[1] = append(Data[1], "testdata2")

	DB := &EZDB.Database{Data: Data, File: "SaveFile.txt"}

	var NewTestSlice []string

	NewTestSlice = append(NewTestSlice, "1")
	NewTestSlice = append(NewTestSlice, "testcol1")
	NewTestSlice = append(NewTestSlice, "testcol2")

	Functions.Replace("1", NewTestSlice, DB)

	if DB.Data[1][0] != "1" {
		test.Errorf("New Column was expected to be 1. Intead got " + DB.Data[1][0])
	}
	if DB.Data[1][1] != "testcol1" {
		test.Errorf("New Column was expected to be testcol1. Intead got " + DB.Data[1][1])
	}
	if DB.Data[1][2] != "testcol2" {
		test.Errorf("New Column was expected to be testcol2. Intead got " + DB.Data[1][2])
	}

}
