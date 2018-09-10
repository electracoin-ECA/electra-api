package models

import (
	"github.com/Electra-project/electra-api/src/database"
	"github.com/globalsign/mgo/bson"
)

const AddrCollectionName = "address"

type Address struct {
	Addr         string `json:"address" bson:"address"`
	Sent         uint
	Received     uint
	Balance      uint
	Transactions []*TransactionSummary `json:"last_txs" bson:"last_txs"`
}

type TransactionSummary struct {
	TransactionHash string `json:"addresses" bson:"addresses"`
	Type            string
}

func StoreAddressIfNotPresent(b *Address) error {
	if present, err := IsAddressPresent(b); err != nil || !present {
		// even in case of an error lets try to insert the address
		return StoreAddress(b)
	}
	// idempotent
	return nil
}

func GetAddressFromDB(hash string) (*Address, error) {
	addr := Address{}
	db := database.Get()
	collection := db.C(AddrCollectionName)
	err := collection.Find(bson.M{
		"address": hash,
	}).One(&addr)
	if err != nil {
		return nil, err
	}
	return &addr, nil
}

func StoreAddress(b *Address) error {
	db := database.Get()
	collection := db.C(AddrCollectionName)
	return collection.Insert(b)
}

func IsAddressPresent(b *Address) (bool, error) {
	db := database.Get()
	collection := db.C(AddrCollectionName)
	count, err := collection.Find(bson.M{
		"address": b.Addr,
	}).Count()

	if err != nil {
		return false, err
	}

	if count == 0 {
		return false, nil
	}

	return true, nil

}
