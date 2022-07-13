package datamodels

import (
	//"strings"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// User is our User example model.
// Keep note that the tags for public-use (for our web app)
// should be kept in other file like "web/viewmodels/user.go"
// which could wrap by embedding the datamodels.User or
// define completely new fields instead but for the sake
// of the example, we will use this datamodel
// as the only one User model in our application.

/*

type User struct {
	// ID             int64     `json:"id" form:"id"`
	ID             string    `json:"id"`
	Firstname      string    `json:"firstname" form:"firstname"`
	Username       string    `json:"username" form:"username"`
	HashedPassword []byte    `json:"-" form:"-"`
	Roles          []Role    `json:"roles"`
	CreatedAt      time.Time `json:"created_at" form:"created_at"`
}
*/

type User struct {
	ID        bson.ObjectId `json:"_id" bson:"_id"`
	Firstname string        `json:"firstname" bson:"firstname" form:"firstname"`
	Username  string        `json:"username" bson:"username" form:"username"`
	Password  []byte        `json:"password" bson:"password" form:"password"`
	Roles     []Role        `json:"roles" bson:"roles"`
	UpdatedAt time.Time     `json:"update_at" bson:"update_at" form:"update_at"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at" form:"created_at"`
}
