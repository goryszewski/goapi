package model

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://root:example@mongo:27017/"

type DB struct {
	client *mongo.Client
}

func (c *DB) Get() *mongo.Collection {
	return c.client.Database("godb").Collection("name")
}

func NewDB(ctx context.Context) *DB {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return &DB{client: client}
}
