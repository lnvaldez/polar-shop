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

func (r *MongoRepo) GetAllAvailableProducts() ([]models.Product, error) {
	collection := r.db.Collection("products")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var products []models.Product
	if err = cursor.All(context.TODO(), &products); err != nil {
		return nil, err
	}
	return products, nil
}