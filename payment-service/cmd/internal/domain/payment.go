package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type PaymentSchema struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	UserId        primitive.ObjectID `bson:"userId,omitempty"`
	Amount        float64            `bson:"amount"`
	PaymentMethod string             `bson:"paymentMethod"`
	Status        string             `bson:"status"`
	TransactionId string             `bson:"transactionId"`
}
