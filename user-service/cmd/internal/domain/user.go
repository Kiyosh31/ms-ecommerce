package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserRole string

const (
	UserRoleUser   UserRole = "USER"
	UserRoleVendor UserRole = "VENDOR"
)

type UserSchema struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty"`
	Name      string               `bson:"name" binding:"required"`
	Email     string               `bson:"email" binding:"required"`
	Password  string               `bson:"password" binding:"required"`
	Addresses []Address            `bson:"addresses"`
	Orders    []primitive.ObjectID `bson:"orders"`
	IsActive  bool                 `bson:"isActive"`
	Role      UserRole             `bson:"role"`
}

type Address struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Name    string             `bson:"name" binding:"required"`
	Street  string             `bson:"street" binding:"required"`
	City    string             `bson:"city" binding:"required"`
	State   string             `bson:"state" binding:"required"`
	Country string             `bson:"country" binding:"required"`
	ZipCode int64              `bson:"zipCode" binding:"required"`
}
