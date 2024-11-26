package db

import (
	"fmt"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"

	"polar-shop/models"
)

type OrderDB struct {
	MongoCollection *mongo.Collection
}

func (db *OrderDB) GetAllOrders() ([]model.Order, error) {
	result, err := db.MongoCollection.Find(context.Background(), bson.D{})

	if err != nil {
		return nil, err
	}

	var orders []model.Order
	err = result.All(context.Background(), &orders)

	if err != nil {
		return nil, fmt.Errorf("error %s", err.Error())
	}

	return orders, nil
}