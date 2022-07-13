package datamodels

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Sale struct {
	ID        bson.ObjectId `json:"_id" bson:"_id"`
	Purchases []Purchase    `json:"purchases" bson:"purchases" form:"purchases"`
	UpdatedAt time.Time     `json:"update_at" bson:"update_at" form:"update_at"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at" form:"created_at"`
}
