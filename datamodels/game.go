package datamodels

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Game struct {
	ID        bson.ObjectId `json:"_id" bson:"_id"`
	Name      string        `json:"name" bson:"name" validate:"required"`
	GameCode  string        `json:"gamecode" bson:"gamecode" validate:"required"`
	UpdatedAt time.Time     `json:"update_at" bson:"update_at" form:"update_at"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at" form:"created_at"`
}
