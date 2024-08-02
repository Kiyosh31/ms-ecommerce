package product_types

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductSchema struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name" binding:"required"`
	Description string             `bson:"description" binding:"required"`
	Price       float32            `bson:"price" binding:"required"`
	Category    CategorySchema     `bson:"category" binding:"required"`
	Brand       BrandSchema        `bson:"brand" binding:"required"`
	Images      []string           `bson:"images" binding:"required"`
	Attributes  Attributes         `bson:"attributes"`
	Inventory   Inventory          `bson:"inventory"`
}

type CategorySchema struct {
	ID             primitive.ObjectID `bson:"_id"`
	Name           string             `bson:"name"`
	Description    string             `bson:"description"`
	ParentCategory primitive.ObjectID `bson:"parentCategory"`
}
type BrandSchema struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
}

type Attributes struct {
	Size     int64  `bson:"size"`
	Color    string `bson:"color"`
	Material string `bson:"material"`
}

type Inventory struct{}

type StockMessageQueue struct {
	ID                string
	ProductId         string
	ProductName       string
	AvailableQuantity int32
}
