package user_types

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserSchema struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `bson:"name" binding:"required"`
	LastName  string             `bson:"lastName" binding:"required"`
	BirthDate string             `bson:"birth" binding:"required"`
	Email     string             `bson:"email" binding:"required"`
	Password  string             `bson:"password" binding:"required"`
	Cards     []Card             `bson:"cards"`
	Addresses []Address          `bson:"addresses"`
}

type Card struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Number         string             `bson:"number" binding:"required"`
	CardHolderName string             `bson:"cardHolderName" binding:"required"`
	Expiration     string             `bson:"expiration" binding:"required"`
	Cvv            int32              `bson:"cvv" binding:"required"`
	CardType       string             `bson:"cardType" binding:"required"`
	IsDefault      bool               `bson:"default" binding:"required"`
}

type Address struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name" binding:"required"`
	Street    string             `bson:"street" binding:"required"`
	City      string             `bson:"city" binding:"required"`
	State     string             `bson:"state" binding:"required"`
	Country   string             `bson:"country" binding:"required"`
	ZipCode   int64              `bson:"zipCode" binding:"required"`
	IsDefault bool               `bson:"default" binding:"required"`
}
