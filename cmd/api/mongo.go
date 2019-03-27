package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoClient returns a mongoDB client from provided url.
func MongoClient(ctx context.Context, url string) (*mongo.Client, error) {
	return mongo.Connect(ctx, options.Client().ApplyURI(url))
}
