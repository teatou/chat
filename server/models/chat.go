package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Chat struct {
	Id            primitive.ObjectID `json:"_id" bson:"_id"`
	Name          string             `json:"name" bson:"name"`
	IsGroup       bool               `json:"is_group" bson:"is_group"`
	Users         []*User            `json:"users" bson:"users"`
	LatestMessage *Message           `json:"latest_message" bson:"latest_message"`
	Admin         *User              `json:"admin" bson:"admin"`
}
