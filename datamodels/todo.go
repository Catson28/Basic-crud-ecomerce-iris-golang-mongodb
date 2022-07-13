package datamodels

import "gopkg.in/mgo.v2/bson"

// Todo represents the Todo model.
type Todo struct {
	ID        bson.ObjectId `json:"_id" bson:"_id"  validate:"required"`
	UserID    string        `json:"user_id"  bson:"user_id"  validate:"required"`
	Title     string        `json:"title"  bson:"title"  validate:"required"`
	Body      string        `json:"body"  bson:"body"  validate:"required"`
	CreatedAt int64         `json:"created_at"  bson:"created_at"` // unix seconds.
}
