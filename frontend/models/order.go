package model 

type Order struct {
	ID string `bson:"_id,omitempty"`
	CustomerName string `bson:"name"`
	Email string `bson:"email"`
	Address string `bson:"address"`
	Status string `bson:"status"`
}