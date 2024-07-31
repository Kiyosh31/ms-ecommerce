package cart_types

import "go.mongodb.org/mongo-driver/bson/primitive"

type CartSchema struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	UserId primitive.ObjectID `bson:"userId" binding:"required"`
	Total  float32            `bson:"totalPrice" binding:"required"`
}
