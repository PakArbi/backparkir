package backparkir

import (
	"encoding/json"
	"net/http"
	// "os"
	"fmt"

	qrcode "github.com/skip2/go-qrcode"
	// "github.com/whatsauth/watoken"
	// "go.mongodb.org/mongo-driver/bson"
)



func GCFCreateParkiran(PUBLICKEY,MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn, err := SetConnection(MONGOCONNSTRINGENV, dbname)
	if err != nil {
		return err.Error()
	}

	var dataparkiran Parkiran
	err = json.NewDecoder(r.Body).Decode(&dataparkiran)
	if err != nil {
		return err.Error()
	}

	result, err := CreateNewParkiran(mconn, collectionname, dataparkiran)
	if err != nil {
		return GCFReturnStruct(CreateResponse(true, fmt.Sprintf("Failed Create Parkiran: %v", err), dataparkiran))
	}

	return GCFReturnStruct(CreateResponse(false, fmt.Sprintf("Success Create Parkiran: %v", result.InsertedID), dataparkiran))
}

// Delete Parkiran
func GCFDeleteParkiran(PUBLICKEY,MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn, err := SetConnection(MONGOCONNSTRINGENV, dbname)
	if err != nil {
		return err.Error()
	}

	var dataparkiran struct {
		ParkiranID int `json:"parkiranid"`
	}
	err = json.NewDecoder(r.Body).Decode(&dataparkiran)
	if err != nil {
		return err.Error()
	}

	result, err := DeleteParkiran(mconn, collectionname, dataparkiran.ParkiranID)
	if err != nil {
		return GCFReturnStruct(CreateResponse(true, fmt.Sprintf("Failed Delete Parkiran: %v", err), dataparkiran))
	}

	if result.DeletedCount == 0 {
		return GCFReturnStruct(CreateResponse(false, "No matching document found to delete", dataparkiran))
	}

	return GCFReturnStruct(CreateResponse(false, "Success Delete Parkiran", dataparkiran))
}

// Update Parkiran
func GCFUpdateParkiran(PUBLICKEY,MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn, err := SetConnection(MONGOCONNSTRINGENV, dbname)
	if err != nil {
		return err.Error()
	}

	var dataparkiran Parkiran
	err = json.NewDecoder(r.Body).Decode(&dataparkiran)
	if err != nil {
		return err.Error()
	}

	result, err := UpdateParkiran(mconn, collectionname, dataparkiran.ParkiranId, dataparkiran)
	if err != nil {
		return GCFReturnStruct(CreateResponse(true, fmt.Sprintf("Failed Update Parkiran: %v", err), dataparkiran))
	}

	if result.ModifiedCount == 0 {
		return GCFReturnStruct(CreateResponse(false, "No matching document found to update", dataparkiran))
	}

	return GCFReturnStruct(CreateResponse(false, "Success Update Parkiran", dataparkiran))
}


// Get All Parkiran
func GCFGetAllParkiran(PUBLICKEY,MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn, err := SetConnection(MONGOCONNSTRINGENV, dbname)
	if err != nil {
		return err.Error()
	}

	dataparkiran, err := GetAllParkiran(mconn, collectionname)
	if err != nil {
		return GCFReturnStruct(CreateResponse(false, fmt.Sprintf("Failed Get All Parkiran: %v", err), dataparkiran))
	}

	return GCFReturnStruct(CreateResponse(true, "Success Get All Parkiran", dataparkiran))
}

func generateCodeQR(parkiran Parkiran) ([]byte, error) {
    // Convert data to JSON
    jsonData, err := json.Marshal(parkiran)
    if err != nil {
        return nil, fmt.Errorf("failed to convert data to JSON: %v", err)
    }

    // Generate QR code
    qrCode, err := qrcode.Encode(string(jsonData), qrcode.Medium, 256)
    if err != nil {
        return nil, fmt.Errorf("failed to generate QR code: %v", err)
    }

    return qrCode, nil
}

// GCFPostParkiran is an example of an HTTP request handler function
func GCFPostParkiran(PUBLICKEY,MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
    mconn, err := SetConnection(MONGOCONNSTRINGENV, dbname)
    if err != nil {
        return err.Error()
    }

    var parkiranData Parkiran // Gantilah "Parkiran" dengan struktur data yang sesuai

    // Mendekode data dari body request menjadi variabel parkiranData
    err = json.NewDecoder(r.Body).Decode(&parkiranData)
    if err != nil {
        return GCFReturnStruct(CreateResponse(false, fmt.Sprintf("Failed to parse parkiran data: %v", err), nil))
    }

    // Memasukkan data parkiran ke dalam database
    err = InsertParkiranData(mconn, collectionname, parkiranData)
    if err != nil {
        return GCFReturnStruct(CreateResponse(false, fmt.Sprintf("Failed to insert parkiran data: %v", err), nil))
    }

    // Generate QR code
    qrCode, err := generateCodeQR(parkiranData)
    if err != nil {
        return GCFReturnStruct(CreateResponse(false, fmt.Sprintf("Failed to generate QR code: %v", err), nil))
    }

    // Simpan QR code ke MongoDB
    err = SaveQRCodeToMongoDB(mconn, "PakArbi", qrCode) // Ganti dengan fungsi yang sesuai
    if err != nil {
        return GCFReturnStruct(CreateResponse(false, fmt.Sprintf("Failed to save QR code to MongoDB: %v", err), nil))
    }

    // Create notification based on Parkiran data
    notification := Notifikasi{
        Status:  200,
        Message: "QR code generated successfully and saved to MongoDB",
        Data:    parkiranData,
    }

    return GCFReturnStruct(CreateResponse(true, "Success inserting parkiran data", notification))
}





// Get All Parkiran By Id
func GCFGetAllParkiranID(PUBLICKEY,MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn, err := SetConnection(MONGOCONNSTRINGENV, dbname)
	if err != nil {
		return err.Error()
	}

	var dataparkiran Parkiran
	err = json.NewDecoder(r.Body).Decode(&dataparkiran)
	if err != nil {
		return err.Error()
	}

	parkiran, err := GetParkiranByID(mconn, collectionname, dataparkiran.ParkiranId)
	if err != nil {
		return GCFReturnStruct(CreateResponse(false, fmt.Sprintf("Failed to Get ID Parkiran: %v", err), dataparkiran))
	}

	if parkiran != nil {
		return GCFReturnStruct(CreateResponse(true, "Success: Get ID Parkiran", parkiran))
	}

	return GCFReturnStruct(CreateResponse(false, "Failed to Get ID Parkiran", dataparkiran))
}


func GCFReturnStruct(DataStuct any) string {
	jsondata, _ := json.Marshal(DataStuct)
	return string(jsondata)
}

func ReturnStringStruct(Data any) string {
	jsonee, _ := json.Marshal(Data)
	return string(jsonee)
}

// <--- FUNCTION PARKIRAN --->
func GCFInsertParkiranNPM(publickey, MONGOCONNSTRINGENV, dbname, colluser, collparkiran string, r *http.Request) string {
	var response Credential
	response.Status = false
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var userdata User
	gettoken := r.Header.Get("Login")
	if gettoken == "" {
		response.Message = "Header Login Not Exist"
	} else {
		// Process the request with the "Login" token
		checktoken := watoken.DecodeGetId(os.Getenv(publickey), gettoken)
		userdata.NPM = checktoken
		if checktoken == "" {
			response.Message = "Kamu kayaknya belum punya akun"
		} else {
			user2 := FindUserNPM(mconn, colluser, userdata)
			if user2.Role == "user" {
				var dataparkiran Parkiran
				err := json.NewDecoder(r.Body).Decode(&dataparkiran)
				if err != nil {
					response.Message = "Error parsing application/json: " + err.Error()
				} else {
					insertParkiran(mconn, collparkiran, Parkiran{
						Parkiranid:     dataparkiran.Parkiranid,
						Nama:           dataparkiran.Nama,
						NPM:            dataparkiran.NPM,
						Prodi:          dataparkiran.Prodi,
						NamaKendaraan:  dataparkiran.NamaKendaraan,
						NomorKendaraan: dataparkiran.NomorKendaraan,
						JenisKendaraan: dataparkiran.JenisKendaraan,
						Status:         dataparkiran.Status,
					})
					response.Status = true
					response.Message = "Berhasil Insert Data Parkiran"
				}
			} else {
				response.Message = "Anda tidak dapat Insert data karena bukan user"
			}
		}
	}
	return GCFReturnStruct(response)
}

func GCFInsertParkiranEmail(publickey, MONGOCONNSTRINGENV, dbname, colluser, collparkiran string, r *http.Request) string {
	var response Credential
	response.Status = false
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var userdata User
	gettoken := r.Header.Get("Login")
	if gettoken == "" {
		response.Message = "Header Login Not Exist"
	} else {
		// Process the request with the "Login" token
		checktoken := watoken.DecodeGetId(os.Getenv(publickey), gettoken)
		userdata.Email = checktoken
		if checktoken == "" {
			response.Message = "Kamu kayaknya belum punya akun"
		} else {
			user2 := FindUserEmail(mconn, colluser, userdata)
			if user2.Role == "user" {
				var dataparkiran Parkiran
				err := json.NewDecoder(r.Body).Decode(&dataparkiran)
				if err != nil {
					response.Message = "Error parsing application/json: " + err.Error()
				} else {
					insertParkiran(mconn, collparkiran, Parkiran{
						Parkiranid:     dataparkiran.Parkiranid,
						Nama:           dataparkiran.Nama,
						NPM:            dataparkiran.NPM,
						Prodi:          dataparkiran.Prodi,
						NamaKendaraan:  dataparkiran.NamaKendaraan,
						NomorKendaraan: dataparkiran.NomorKendaraan,
						JenisKendaraan: dataparkiran.JenisKendaraan,
						Status:         dataparkiran.Status,
					})
					response.Status = true
					response.Message = "Berhasil Insert Data Parkiran"
				}
			} else {
				response.Message = "Anda tidak dapat Insert data karena bukan user"
			}
		}
	}
	return GCFReturnStruct(response)
}

func GCFUpdateParkiranNPM(publickey, MONGOCONNSTRINGENV, dbname, colluser, collparkiran string, r *http.Request) string {
	var response Credential
	response.Status = false
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var userdata User

	gettoken := r.Header.Get("Login")
	if gettoken == "" {
		response.Message = "Header Login Not Exist"
	} else {
		checktoken := watoken.DecodeGetId(os.Getenv(publickey), gettoken)
		userdata.NPM = checktoken
		if checktoken == "" {
			response.Message = "Kamu kayaknya belum punya akun"
		} else {
			user2 := FindUserNPM(mconn, colluser, userdata)
			if user2.Role == "user" {
				var dataparkiran Parkiran
				err := json.NewDecoder(r.Body).Decode(&dataparkiran)
				if err != nil {
					response.Message = "Error parsing application/json: " + err.Error()
				} else {
					UpdatedParkiran(mconn, collparkiran, bson.M{"id": dataparkiran.ID}, dataparkiran)
					response.Status = true
					response.Message = "Berhasil Update Parkiran"
					GCFReturnStruct(CreateResponse(true, "Success Update Parkiran", dataparkiran))
				}
			} else {
				response.Message = "Anda tidak dapat Update data karena bukan user"
			}
		}
	}
	return GCFReturnStruct(response)
}

func GCFUpdateParkiranEmail(publickey, MONGOCONNSTRINGENV, dbname, colluser, collparkiran string, r *http.Request) string {
	var response Credential
	response.Status = false
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var userdata User

	gettoken := r.Header.Get("Login")
	if gettoken == "" {
		response.Message = "Header Login Not Exist"
	} else {
		checktoken := watoken.DecodeGetId(os.Getenv(publickey), gettoken)
		userdata.Email = checktoken
		if checktoken == "" {
			response.Message = "Kamu kayaknya belum punya akun"
		} else {
			user2 := FindUserEmail(mconn, colluser, userdata)
			if user2.Role == "user" {
				var dataparkiran Parkiran
				err := json.NewDecoder(r.Body).Decode(&dataparkiran)
				if err != nil {
					response.Message = "Error parsing application/json: " + err.Error()
				} else {
					UpdatedParkiran(mconn, collparkiran, bson.M{"id": dataparkiran.ID}, dataparkiran)
					response.Status = true
					response.Message = "Berhasil Update Parkiran"
					GCFReturnStruct(CreateResponse(true, "Success Update Parkiran", dataparkiran))
				}
			} else {
				response.Message = "Anda tidak dapat Update data karena bukan user"
			}
		}
	}
	return GCFReturnStruct(response)
}

func GCFDeleteParkiranNPM(publickey, MONGOCONNSTRINGENV, dbname, colluser, collparkiran string, r *http.Request) string {
	var respon Credential
	respon.Status = false
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var userdata User

	gettoken := r.Header.Get("Login")
	if gettoken == "" {
		respon.Message = "Header Login Not Exist"
	} else {
		// Process the request with the "Login" token
		checktoken := watoken.DecodeGetId(os.Getenv(publickey), gettoken)
		userdata.NPM = checktoken
		if checktoken == "" {
			respon.Message = "Kamu kayaknya belum punya akun"
		} else {
			user2 := FindUserNPM(mconn, colluser, userdata)
			if user2.Role == "user" {
				var dataparkiran Parkiran
				err := json.NewDecoder(r.Body).Decode(&dataparkiran)
				if err != nil {
					respon.Message = "Error parsing application/json: " + err.Error()
				} else {
					DeleteParkiran(mconn, collparkiran, dataparkiran)
					respon.Status = true
					respon.Message = "Berhasil Delete Parkiran"
				}
			} else {
				respon.Message = "Anda tidak dapat Delete data karena bukan user"
			}
		}
	}
	return GCFReturnStruct(respon)
}

func GCFDeleteParkiranEmail(publickey, MONGOCONNSTRINGENV, dbname, colluser, collparkiran string, r *http.Request) string {
	var respon Credential
	respon.Status = false
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var userdata User

	gettoken := r.Header.Get("Login")
	if gettoken == "" {
		respon.Message = "Header Login Not Exist"
	} else {
		// Process the request with the "Login" token
		checktoken := watoken.DecodeGetId(os.Getenv(publickey), gettoken)
		userdata.Email = checktoken
		if checktoken == "" {
			respon.Message = "Kamu kayaknya belum punya akun"
		} else {
			user2 := FindUserEmail(mconn, colluser, userdata)
			if user2.Role == "user" {
				var dataparkiran Parkiran
				err := json.NewDecoder(r.Body).Decode(&dataparkiran)
				if err != nil {
					respon.Message = "Error parsing application/json: " + err.Error()
				} else {
					DeleteParkiran(mconn, collparkiran, dataparkiran)
					respon.Status = true
					respon.Message = "Berhasil Delete Parkiran"
				}
			} else {
				respon.Message = "Anda tidak dapat Delete data karena bukan user"
			}
		}
	}
	return GCFReturnStruct(respon)
}

func GetAllDataParkiran(PublicKey, MongoEnv, dbname, colname string, r *http.Request) string {
	req := new(Response)
	conn := SetConnection(MongoEnv, dbname)
	tokenlogin := r.Header.Get("Login")
	if tokenlogin == "" {
		req.Status = false
		req.Message = "Header Login Not Found"
	} else {
		// Dekode token untuk mendapatkan
		_, err := DecodeGetParkiran(os.Getenv(PublicKey), tokenlogin)
		if err != nil {
			req.Status = false
			req.Message = "Data Tersebut tidak ada" + tokenlogin
		} else {
			// Langsung ambil data catalog
			dataparkiran := GetAllParkiran(conn, colname)
			if dataparkiran == nil {
				req.Status = false
				req.Message = "Data Parkiran tidak ada"
			} else {
				req.Status = true
				req.Message = "Data Parkiran berhasil diambil"
				req.Data = dataparkiran
			}
		}
	}
	return ReturnStringStruct(req)
}

func GetOneDataParkiran(PublicKey, MongoEnv, dbname, colname string, r *http.Request) string {
	req := new(ResponseParkiran)
	resp := new(RequestParkiran)
	conn := MongoCreateConnection(MongoEnv, dbname)
	tokenlogin := r.Header.Get("Login")
	if tokenlogin == "" {
		req.Status = false
		req.Message = "Header Login Not Found"
	} else {
		err := json.NewDecoder(r.Body).Decode(&resp)
		if err != nil {
			req.Message = "error parsing application/json: " + err.Error()
		} else {
			dataparkiran := GetOneParkiranData(conn, colname, resp.Parkiranid)
			req.Status = true
			req.Message = "data Parkiran berhasil diambil"
			req.Data = dataparkiran
		}
	}
	return ReturnStringStruct(req)
}

func GCFGetAllParkiranID(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)

	var dataparkiran Parkiran
	err := json.NewDecoder(r.Body).Decode(&dataparkiran)
	if err != nil {
		return err.Error()
	}

	parkiran := GetAllParkiranID(mconn, collectionname, dataparkiran)
	if parkiran != (Parkiran{}) {
		return GCFReturnStruct(CreateResponse(true, "Success: Get ID Parkiran", dataparkiran))
	} else {
		return GCFReturnStruct(CreateResponse(false, "Failed to Get ID Parkiran", dataparkiran))
	}
}
