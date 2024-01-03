package backparkir

import (
	"fmt"
	"testing"
)


func TestParkiran(t *testing.T) {
	mconn, err := SetConnection("MONGOSTRING", "PakArbi")
	if err != nil {
		t.Fatalf("Error connecting to MongoDB: %v", err)
	}

	var parkirandata = Parkiran{
		ParkiranId:     1,
		Nama:           "Muhammad Faisal Ash",
		NPM:            "1214020",
		Jurusan:        "D4 Teknik Informatika",
		NamaKendaraan:  "Supra X 125",
		NomorKendaraan: "F 1234 NR",
		JenisKendaraan: "Motor",
	}

	result, err := CreateNewParkiran(mconn, "Parkiran", parkirandata)
	if err != nil {
		t.Fatalf("Error creating parkiran: %v", err)
	}

	fmt.Printf("InsertedID: %v\n", result.InsertedID)
}

func TestAllParkiran(t *testing.T) {
	mconn, err := SetConnection("MONGOSTRING", "PakArbi")
	if err != nil {
		t.Fatalf("Error connecting to MongoDB: %v", err)
	}

	parkiran, err := GetAllParkiran(mconn, "Parkiran")
	if err != nil {
		t.Fatalf("Error fetching all parkiran: %v", err)
	}

	for _, p := range parkiran {
		fmt.Println(p)
	}
}