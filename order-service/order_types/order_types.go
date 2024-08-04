package ordertypes

import "go.mongodb.org/mongo-driver/bson/primitive"

type OrderSchema struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
}
