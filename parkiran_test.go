package backparkir

import (
	"fmt"
	"testing"

	"github.com/aiteung/atdb"
	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson"
)


func TestParkiran(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "PakArbi")
	var parkirandata Parkiran
	parkirandata.ParkiranId = "1"
	parkirandata.Nama = "Farhan Rizki Maulana"
	parkirandata.NPM = "1214020"
	parkirandata.Prodi = "D4 Teknik Informatika"
	parkirandata.NamaKendaraan = "Supra X 125"
	parkirandata.NomorKendaraan = "F 1234 NR"
	parkirandata.JenisKendaraan = "Motor"
	CreateNewParkiran(mconn, "parkiran", parkirandata)
}

func TestAllParkiran(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "PakArbi")
	parkiran := GetAllParkiran(mconn, "parkiran")
	fmt.Println(parkiran)
}

// func TestParkiran(t *testing.T) {
// 	mconn, err := SetConnection("MONGOSTRING", "PakArbi")
// 	if err != nil {
// 		t.Fatalf("Error connecting to MongoDB: %v", err)
// 	}

// 	var parkirandata = Parkiran{
// 		ParkiranId:     1,
// 		Nama:           "Muhammad Faisal Ashshidiq",
// 		NPM:            "1214041",
// 		Jurusan:        "D4 Teknik Informatika",
// 		NamaKendaraan:  "Mio Z",
// 		NomorKendaraan: "D 3316 GXF",
// 		JenisKendaraan: "Motor",
// 	}

// 	result, err := CreateNewParkiran(mconn, "Parkiran", parkirandata)
// 	if err != nil {
// 		t.Fatalf("Error creating parkiran: %v", err)
// 	}

// 	fmt.Printf("InsertedID: %v\n", result.InsertedID)
// }

// func TestAllParkiran(t *testing.T) {
// 	mconn, err := SetConnection("MONGOSTRING", "PakArbi")
// 	if err != nil {
// 		t.Fatalf("Error connecting to MongoDB: %v", err)
// 	}

// 	parkiran, err := GetAllParkiran(mconn, "Parkiran")
// 	if err != nil {
// 		t.Fatalf("Error fetching all parkiran: %v", err)
// 	}

// 	for _, p := range parkiran {
// 		fmt.Println(p)
// 	}
// }


// func TestGenerateQRCode(t *testing.T) {
// 	formData := FormData{
// 		NamaLengkap:    "Farhan Rizki Maulana",
// 		NPM:            "1214020",
// 		Jurusan:        "D4 Teknik Informatika",
// 		NamaKendaraan:  "Supra X 125",
// 		NomorKendaraan: "F 1234 NR",
// 		JenisKendaraan: "Motor",
// 	}

// 	err := GenerateQRCode(formData)
// 	if err != nil {
// 		t.Errorf("Failed to generate QR code: %v", err)
// 	}

// 	// Check if QR code file exists
// 	if _, err := os.Stat("qrcode.png"); os.IsNotExist(err) {
// 		t.Errorf("QR code file does not exist: %v", err)
// 	}

// 	// Check if JSON data is generated correctly
// 	dataJSON, _ := json.Marshal(formData)

// 	expectedJSON := `{"namalengkap":"Farhan Rizki Maulana","npm":"1214020","jurusan":"D4 Teknik Informatika","namakendaraan":"Supra X 125","nomorkendaraan":"F 1234 NR","jeniskendaraan":"Motor"}`

// 	// Validate JSON data
// 	if string(dataJSON) != expectedJSON {
// 		t.Errorf("Incorrect JSON data generated")
// 	}
// }