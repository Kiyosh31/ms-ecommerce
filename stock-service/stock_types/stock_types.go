package stocktypes

import "go.mongodb.org/mongo-driver/bson/primitive"

type StockSchema struct {
	ID                primitive.ObjectID `bson:"_id"`
	ProductId         primitive.ObjectID `bson:"productId"`
	ProductName       string             `bson:"productName"`
	AvailableQuantity int32              `bson:"quantity"`
}
