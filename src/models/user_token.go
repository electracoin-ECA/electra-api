package models

import (
	"time"

	"github.com/Electra-project/electra-api/src/database"
	"github.com/Electra-project/electra-api/src/libs/mnemonic"
	"github.com/globalsign/mgo/bson"
)

// UserToken model.
type UserToken struct {
	ID        bson.ObjectId `bson:"_id" json:"-"`
	Challenge string        `bson:"challenge" json:"challenge"`
	PurseHash string        `bson:"purseHash" json:"purseHash"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time     `bson:"updatedAt" json:"updatedAt"`
}

// GetByPurseHash finds a user token in the database
// by its Purse Account address hash.
func (u UserToken) GetByPurseHash(purseHash string) (*UserToken, error) {
	db := database.Get()
	collection := db.C("users-tokens")

	var userToken *UserToken
	err := collection.Find(bson.M{"purseHash": purseHash}).One(&userToken)
	if err != nil {
		return nil, err
	}

	return userToken, nil
}

// Insert generates and creates a new user token in the database.
func (u UserToken) Insert(purseHash string) (*UserToken, error) {
	db := database.Get()
	collection := db.C("users-tokens")

	challenge, err := generateMnemonic()
	if err != nil {
		return nil, err
	}

	err = collection.Insert(bson.M{
		"challenge": challenge,
		"purseHash": purseHash,
		"createdAt": time.Now(),
		"updatedAt": time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return u.GetByPurseHash(purseHash)
}

func generateMnemonic() (string, error) {
	entropy, err := mnemonic.NewEntropy(256)
	if err != nil {
		return "", err
	}

	mnemonic, err := mnemonic.NewMnemonic(entropy)
	if err != nil {
		return "", err
	}

	return mnemonic, nil
}
