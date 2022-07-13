package datamodels

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Address struct {
	ID        bson.ObjectId `json:"_id" bson:"_id"`
	Casa      string        `json:"casa" bson:"casa" form:"casa"`
	Cidade    float64       `json:"cidade" bson:"cidade" form:"cidade"`
	Contact   float64       `json:"contact" bson:"contact" form:"contact"`
	Pais      string        `json:"pais" bson:"pais" form:"pais"`
	UpdatedAt time.Time     `json:"update_at" bson:"update_at" form:"update_at"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at" form:"created_at"`
}
