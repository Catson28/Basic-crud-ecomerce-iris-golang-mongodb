package datamodels

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Purchase struct {
	ID        bson.ObjectId `json:"_id" bson:"_id"`
	Quantity  int64         `json:"quantity" bson:"quantity" form:"quantity"`
	ProductId bson.ObjectId `json:"productId" bson:"productId" form:"productId"`
	UpdatedAt time.Time     `json:"update_at" bson:"update_at" form:"update_at"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at" form:"created_at"`
}
