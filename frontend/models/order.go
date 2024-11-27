package model 

type Order struct {
	ID string `bson:"_id,omitempty"`
	CustomerName string `bson:"customerName"`
	Email string `bson:"email"`
	Product string `bson:"productName"`
	ProductQuantity int32 `bson:"productQuantity"`
	Address string `bson:"address"`
	Status string `bson:"status"`
}