package inventory_types

import "go.mongodb.org/mongo-driver/bson/primitive"

type InventorySchema struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ProductID primitive.ObjectID `bson:"productId,omitempty"`
	Quantity  int64              `bson:"quantity"`
	Location  string             `bson:"location"`
}
