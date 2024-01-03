package backparkir

import (
	"context"
	"fmt"
	"os"

	"github.com/aiteung/atdb"
	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// <--- FUNCTION CRUD PARKIRAN --->

// parkiran
func CreateNewParkiran(mongoconn *mongo.Database, collection string, parkirandata Parkiran) interface{} {
	return atdb.InsertOneDoc(mongoconn, collection, parkirandata)
}

// parkiran function
func insertParkiran(mongoconn *mongo.Database, collection string, parkirandata Parkiran) interface{} {
	return atdb.InsertOneDoc(mongoconn, collection, parkirandata)
}

func DeleteParkiran(mongoconn *mongo.Database, collection string, parkirandata Parkiran) interface{} {
	filter := bson.M{"parkiranid": parkirandata.Parkiranid}
	return atdb.DeleteOneDoc(mongoconn, collection, filter)
}

func UpdatedParkiran(mongoconn *mongo.Database, collection string, filter bson.M, parkirandata Parkiran) interface{} {
	updatedFilter := bson.M{"parkiranid": parkirandata.Parkiranid}
	return atdb.ReplaceOneDoc(mongoconn, collection, updatedFilter, parkirandata)
}

func GetAllParkiran(mongoconn *mongo.Database, collection string) []Parkiran {
	parkiran := atdb.GetAllDoc[[]Parkiran](mongoconn, collection)
	return parkiran
}

func GetOneParkiran(mongoconn *mongo.Database, collection string, parkirandata Parkiran) interface{} {
	filter := bson.M{"parkiranid": parkirandata.Parkiranid}
	return atdb.GetOneDoc[Parkiran](mongoconn, collection, filter)
}

func GetAllParkiranID(mongoconn *mongo.Database, collection string, parkirandata Parkiran) Parkiran {
	filter := bson.M{
		"parkiranid":     parkirandata.Parkiranid,
		"nama":           parkirandata.Nama,
		"npm":            parkirandata.NPM,
		"prodi":          parkirandata.Prodi,
		"namakendaraan":  parkirandata.NamaKendaraan,
		"nomorkendaraan": parkirandata.NomorKendaraan,
		"jeniskendaraan": parkirandata.JenisKendaraan,
	}
	parkiranID := atdb.GetOneDoc[Parkiran](mongoconn, collection, filter)
	return parkiranID
}