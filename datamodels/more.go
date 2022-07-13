package datamodels

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type More struct {
	ID        bson.ObjectId `json:"_id" bson:"_id"`
	Name      string        `json:"name" bson:"name" form:"name"`
	Password  []byte        `json:"password" bson:"password" form:"password"`
	Games     []string      `json:"games" bson:"games"`
	UpdatedAt time.Time     `json:"update_at" bson:"update_at" form:"update_at"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at" form:"created_at"`
}
