package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserSchema struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name" binding:"required"`
	LastName string             `bson:"lastName" binding:"required"`
	Birth    string             `bson:"birth" binding:"required"`
	Email    string             `bson:"email" binding:"required"`
	Password string             `bson:"password" binding:"required"`
	Card     []Card             `bson:"card" binding:"required"`
	Address  []Address          `bson:"address"`
}

type Card struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Number     int64              `bson:"number"`
	Cvv        int64              `bson:"cvv"`
	Expiration string             `bson:"expiration"`
	Default    bool               `bson:"default"`
}

type Address struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Name    string             `bson:"name"`
	ZipCode int64              `bson:"zipCode"`
	Default bool               `bson:"default"`
}
