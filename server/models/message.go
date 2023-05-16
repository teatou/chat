package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	Id      primitive.ObjectID `json:"_id" bson:"_id"`
	From    *User              `json:"from" bson:"from"`
	Content string             `json:"content" bson:"content"`
	Chat    *Chat              `json:"chat" bson:"chat"`
}
