package backparkir

import "go.mongodb.org/mongo-driver/bson/primitive"

type ResponseBack struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    []string `json:"data"`
}

type ResponseParkiran struct {
	Status  bool     `json:"status"`
	Message string   `json:"message"`
	Data    Parkiran `json:"data"`
}

type RequestParkiran struct {
	Parkiranid string `json:"parkiranid"`
}

type Notifikasi struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    Parkiran  	`json:"data"`
}

type Time struct {
	Message     string `json:"message,omitempty" bson:"message,omitempty"`
	WaktuMasuk  string `json:"waktumasuk,omitempty" bson:"waktumasuk,omitempty"`
	WaktuKeluar string `json:"waktukeluar,omitempty" bson:"waktukeluar,omitempty"`
}

type Parkiran struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" `
	ParkiranId     int                `json:"parkiranid" bson:"parkiranid"`
	Nama           string             `json:"nama" bson:"nama"`
	NPM            string             `json:"npm" bson:"npm"`
	Prodi        string             `json:"prodi" bson:"prodi"`
	NamaKendaraan  string             `json:"namakendaraan" bson:"namakendaraan"`
	NomorKendaraan string             `bson:"nomorkendaraan,omitempty" json:"nomorkendaraan,omitempty"`
	JenisKendaraan string             `json:"jeniskendaraan,omitempty" bson:"jeniskendaraan,omitempty"`
}

type Credential struct {
	Status  bool   `json:"status" bson:"status"`
	Token   string `json:"token,omitempty" bson:"token,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

type Response struct {
	Status  bool        `json:"status" bson:"status"`
	Message string      `json:"message" bson:"message"`
	Data    interface{} `json:"data" bson:"data"`
}

type Prodi struct {
	ProdiId     int 	`json:"prodiid" bson:"prodiid,omitempty"`
	ProdiName    string 	`json:"prodiname" bson:"prodiname,omitempty"`
}


type Updated struct {
	Message string `json:"message"`
}

// type RequestParkiran struct {
// 	Username string `json:"username" bson:"username,omitempty"`
// 	Npm      string `json:"npm" bson:"npm,omitempty"`
// 	Email    string `json:"email" bson:"email,omitempty"`
// 	Message  string `json:"message,omitempty" bson:"message,omitempty"`
// }
