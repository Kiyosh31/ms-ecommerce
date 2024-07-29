package product_types

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductSchema struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	SellerID    primitive.ObjectID `bson:"sellerId,omitempty"`
	Name        string             `bson:"name" binding:"required"`
	Price       float32            `bson:"price" binding:"required"`
	Description string             `bson:"description" binding:"required"`
}
