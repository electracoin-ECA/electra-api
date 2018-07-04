package models

import (
	"github.com/Electra-project/electra-api/src/database"
)

const TransactionCollectionName = "txns"

type Transaction struct {
	TxID       string `json:"txid" bson:"txid"`
	Blockhash  string `json:"blockhash" bson:"blockhash"`
	Blockindex uint   `json:"blockindex" bson:"blockindex"`
	Timestamp  uint   `json:"timestamp" bson:"timestamp"`
	Total      uint   `json:"total" bson:"total"`
	Vout       []*V   `json:"vout" bson:"vout"`
	Vin        []*V   `json:"vin" bson:"vin"`
}

type V struct {
	Address string `json:"addresses" bson:"addresses"`
	Amount  uint   `json:"amount" bson:"amount"`
}

func StoreTransaction(txn *Transaction) error {
	db := database.Get()
	collection := db.C(TransactionCollectionName)
	return collection.Insert(txn)
}
