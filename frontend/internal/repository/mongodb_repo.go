package repository 

import (
	"context"
	"frontend/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type MongoRepo struct {
	client *mongo.Client 
	db *mongo.Database
}

func NewMongoRepo(client *mongo.Client, dbName string) *MongoRepo {
	return &MongoRepo{
		client: client,
		db: client.Database(dbName),
	}
}