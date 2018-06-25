package models

import (
	"log"
	"time"

	"github.com/Electra-project/electra-auth/src/database"
	"github.com/Electra-project/electra-auth/src/helpers"
	"github.com/globalsign/mgo/bson"
)

// User model.
type User struct {
	ID               bson.ObjectId   `bson:"_id" json:"-"`
	PurseHash        string          `bson:"purseHash" json:"purseHash"`
	Token            string          `bson:"token" json:"token"`
	HasAutoUpdate    bool            `bson:"hasAutoUpdate" json:"hasAutoUpdate"`
	IsSynchronized   bool            `bson:"isSynchronized" json:"isSynchronized"`
	PursePrivateKey  string          `bson:"pursePrivateKey" json:"-"`
	TwitterUsername  string          `bson:"twitterUsername" json:"twitterUsername"`
	TwitterCheckedAt time.Time       `bson:"twitterCheckedAt" json:"twitterCheckedAt"`
	BootstrapNodes   []bson.ObjectId `bson:"bootstrapNodes" json:"-"`
	CreatedAt        time.Time       `bson:"createdAt" json:"createdAt"`
	UpdatedAt        time.Time       `bson:"updatedAt" json:"updatedAt"`
}

// UserEditable model.
type UserEditable struct {
	HasAutoUpdate   bool   `bson:"hasAutoUpdate" json:"hasAutoUpdate"`
	IsSynchronized  bool   `bson:"isSynchronized" json:"isSynchronized"`
	TwitterUsername string `bson:"twitterUsername" json:"twitterUsername"`
}

const tokenLength uint8 = 196

// GetByPurseHash gets a user from the database
// by their Purse Account address hash.
func (u User) GetByPurseHash(purseHash string) (*User, error) {
	db := database.Get()
	collection := db.C("users")

	var user *User
	err := collection.Find(bson.M{"purseHash": purseHash}).One(&user)
	if err != nil {
		log.Println(err)

		return nil, err
	}

	return user, nil
}

// Insert creates a new user in the database.
func (u User) Insert(purseHash string) (*User, error) {
	token, err := helpers.GenerateToken(tokenLength)
	if err != nil {
		return nil, err
	}

	db := database.Get()
	collection := db.C("users")

	err = collection.Insert(bson.M{
		"purseHash":       purseHash,
		"token":           token,
		"hasAutoUpdate":   true,
		"isSynchronized":  true,
		"twitterUsername": "",
		"createdAt":       time.Now(),
		"updatedAt":       time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return u.GetByPurseHash(purseHash)
}

// Update the data of an exising user in the database.
func (u User) Update(purseHash string, data UserEditable) (*User, error) {
	db := database.Get()
	collection := db.C("users")

	err := collection.Update(
		bson.M{"purseHash": purseHash},
		bson.M{"$set": bson.M{
			"hasAutoUpdate":   data.HasAutoUpdate,
			"isSynchronized":  data.IsSynchronized,
			"twitterUsername": data.TwitterUsername,
			"updatedAt":       time.Now(),
		}},
	)
	if err != nil {
		return nil, err
	}

	return u.GetByPurseHash(purseHash)
}

// UpdatePursePrivateKey for the user owning this purse address.
func (u User) UpdatePursePrivateKey(
	purseHash string,
	pursePrivateKey string,
) (*User, error) {
	db := database.Get()
	collection := db.C("users")

	err := collection.Update(
		bson.M{"purseHash": purseHash},
		bson.M{"$set": bson.M{
			"pursePrivateKey": pursePrivateKey,
			"updatedAt":       time.Now(),
		}},
	)
	if err != nil {
		return nil, err
	}

	return u.GetByPurseHash(purseHash)
}
