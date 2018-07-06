package models

import (
	"github.com/Electra-project/electra-api/src/database"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

const BlockCollectionName = "blocks"

type Block struct {
	Hash              string  `bson:"hash" json:"hash"`
	Confirmations     uint    `bson:"confirmation" json:"confirmation"`
	Size              uint    `bson:"size" json:"size"`
	Height            uint    `bson:"height" json:"height"`
	Version           uint    `bson:"version" json:"version"`
	MerkeleRoot       string  `bson:"merkeleroot" json:"merkeleroot"`
	Mint              float64 `bson:"mint" json:"mint"`
	Time              uint    `bson:"time" json:"time"`
	Difficulty        float64 `bson:"difficulty" json:"difficulty"`
	Blocktrust        string  `bson:"blocktrust" json:"blocktrust"`
	Chaintrust        string  `bson:"chaintrust" json:"chaintrust"`
	Nextblockhash     string  `bson:"nextblockhash" json:"nextblockhash"`
	Previousblockhash string  `bson:"previousblockhash" json:"previousblockhash"`
	EntropyBit        int     `bson:"entropybit" json:"entropybit"`
	Modifier          string  `bson:"modifier" json:"modifier"`
	Modifierchecksum  string  `bson:"modifierchecksum" json:"modifierchecksum"`
	Signature         string  `bson:"signature" json:"signature"`
}

func StoreBlock(b *Block) error {
	db := database.Get()
	collection := db.C(BlockCollectionName)
	return collection.Insert(b)
}

// IsBlockPresent checks if the block with the hash provided is present or not
func IsBlockPresent(blockHash string) (bool, error) {
	count, err := blockQuery(blockHash).Count()
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil

}

func blockQuery(blockHash string) *mgo.Query {
	db := database.Get()
	collection := db.C(BlockCollectionName)
	return collection.Find(bson.M{
		"hash": blockHash,
	})
}
