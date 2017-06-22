package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Todo struct {
	Id    bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string        `json:"title" binding:"required" bson:"title"`
	Done  int           `json:"done" binding:"required" bson:"done"`
}