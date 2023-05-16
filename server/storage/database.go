package storage

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var UserCollection *mongo.Collection
var ChatCollection *mongo.Collection
var MessageCollection *mongo.Collection

func ConnectToDb() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO")))
	if err != nil {
		log.Fatal("Error connecting to db")
	}

	UserCollection = client.Database("chat").Collection("users")
	ChatCollection = client.Database("chat").Collection("chats")
	MessageCollection = client.Database("chat").Collection("messages")
}
