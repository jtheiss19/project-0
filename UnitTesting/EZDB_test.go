package tests

import (
	"testing"

	EZDB "github.com/jtheiss19/project-0/Database"
)

func TestGrabDBCol(test *testing.T) {
	DB := EZDB.ReadDB("Data1.txt")
	DBNew := DB.GrabDBCol("testing1")

	DBNew.PrettyPrint()

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

	if DBNew != DB {
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
