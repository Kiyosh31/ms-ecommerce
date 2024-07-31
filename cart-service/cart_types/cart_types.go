package cart_types

import "go.mongodb.org/mongo-driver/bson/primitive"

type CartSchema struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	UserId   primitive.ObjectID `bson:"userId" binding:"required"`
	Products []Product          `bson:"products"`
	Total    float32            `bson:"totalPrice" binding:"required"`
}

type Product struct {
	ProductId string  `bson:"productId"`
	Quantity  int64   `bson:"quantity"`
	Price     float64 `bson:"price"`
}
