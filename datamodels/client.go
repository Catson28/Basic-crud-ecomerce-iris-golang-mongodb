package datamodels

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Client struct {
	ID        bson.ObjectId `json:"_id" bson:"_id"`
	Name      string        `json:"name" bson:"name" form:"name"`
	Contact   float64       `json:"contact" bson:"contact" form:"contact"`
	Address   bson.ObjectId `json:"address" bson:"address" form:"address"`
	UpdatedAt time.Time     `json:"update_at" bson:"update_at" form:"update_at"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at" form:"created_at"`
}
