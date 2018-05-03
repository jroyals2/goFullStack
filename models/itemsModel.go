package models

import (
	"gopkg.in/mgo.v2/bson"
)

//Item is the basic model for the database
type Item struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Description string        `bson:"description" json:"description"`
}
