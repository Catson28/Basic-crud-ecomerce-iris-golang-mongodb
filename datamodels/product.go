package datamodels

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Product struct {
	ID        bson.ObjectId `json:"_id" bson:"_id"`
	Name      string        `json:"name" bson:"name" form:"name"`
	Price     float64       `json:"price" bson:"price" form:"price"`
	Catid     bson.ObjectId `json:"catid" bson:"catid" form:"catid"`
	Cost      float64       `json:"cost" bson:"cost" form:"cost"`
	Barcod    float64       `json:"barcod" bson:"barcod" form:"barcod"`
	UpdatedAt time.Time     `json:"update_at" bson:"update_at" form:"update_at"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at" form:"created_at"`
}
