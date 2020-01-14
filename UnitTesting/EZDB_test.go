package tests

import (
	"testing"

	EZDB "github.com/jtheiss19/project-0/Database"
)

func TestGrabDBCol(test *testing.T) {
	DB := EZDB.ReadDB("Data1.txt")
	DBNew := DB.GrabDBCol("testing1")

	if DBNew.Data[0][0] != "Key" {
		test.Errorf("New Column was expected to be Key. Intead got " + DBNew.Data[0][len(DB.Data[0])-1])
	}
	if DBNew.Data[0][1] != "testing1" {
		test.Errorf("New Column was expected to be testing1. Intead got " + DBNew.Data[0][len(DB.Data[0])-1])
	}
	if DBNew.Data[1][0] != "1" {
		test.Errorf("New Column was expected to be 1. Intead got " + DBNew.Data[0][len(DB.Data[0])-1])
	}
	if DBNew.Data[1][1] != "testdata1" {
		test.Errorf("New Column was expected to be testdata1. Intead got " + DBNew.Data[0][len(DB.Data[0])-1])
	}
}

func TestGrabDBRow(test *testing.T) {
	DB := EZDB.ReadDB("Data1.txt")
	DBNew := DB.GrabDBRow(1)

	if DBNew.Data[1][0] != DB.Data[1][0] {
		test.Errorf("New DB was expected to Match the old one")
	}
	if DBNew.Data[1][1] != DB.Data[1][1] {
		test.Errorf("New DB was expected to Match the old one")
	}
	if DBNew.Data[1][2] != DB.Data[1][2] {
		test.Errorf("New DB was expected to Match the old one")
	}
}

func TestGetRowKey(test *testing.T) {
	DB := EZDB.ReadDB("Data1.txt")
	KeyID := DB.GetRowKey("1")

	if KeyID != 1 {
		test.Errorf("Was expecting 1 instead got: " + string(KeyID))
	}
}

func TestAddRow(test *testing.T) {
	DB := EZDB.ReadDB("Data1.txt")

	var Data []string

	Data = append(Data, "2")
	Data = append(Data, "testcol1")
	Data = append(Data, "testcol2")

	DB.AddRow(Data)

	if DB.Data[2][0] != "2" {
		test.Errorf("New Column was expected to be Key. Intead got " + DB.Data[0][len(DB.Data[0])-1])
	}
	if DB.Data[2][1] != "testcol1" {
		test.Errorf("New Column was expected to be testing1. Intead got " + DB.Data[0][len(DB.Data[0])-1])
	}
	if DB.Data[2][2] != "testcol2" {
		test.Errorf("New Column was expected to be 1. Intead got " + DB.Data[0][len(DB.Data[0])-1])
	}

}

func TestDelRow(test *testing.T) {
	DB := EZDB.ReadDB("Data1.txt")

	DB.DelRow(1)

	if len(DB.Data) != 1 {
		test.Errorf("New Column was expected to be 1. Intead got " + DB.Data[0][len(DB.Data[0])-1])
	}

}

func TestCreateCol(test *testing.T) {
	DB := EZDB.ReadDB("Data1.txt")

	DB.CreateCol("Testing")

	if DB.Data[0][len(DB.Data[0])-1] != "Testing" {
		test.Errorf("New Column was expected to be 1. Intead got " + DB.Data[0][len(DB.Data[0])-1])
	}
}

func TestDelCol(test *testing.T) {
	DB := EZDB.ReadDB("Data1.txt")

	DB.DelCol("testing2")

	if DB.Data[0][len(DB.Data[0])-1] == "testing2" {
		test.Errorf("New Column was expected to be 1. Intead got " + DB.Data[0][len(DB.Data[0])-1])
	}

}
