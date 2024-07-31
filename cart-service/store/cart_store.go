package store

import (
	"context"

	"github.com/Kiyosh31/ms-ecommerce/cart-service/cart_types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CartStore struct {
	client              *mongo.Client
	database_name       string
	database_collection string
}

func NewCartStore(client *mongo.Client,
	database_name string,
	database_collection string) *CartStore {
	return &CartStore{
		client:              client,
		database_name:       database_name,
		database_collection: database_collection,
	}
}

func (s *CartStore) getCartCollection() *mongo.Collection {
	return s.client.Database(s.database_name).Collection(s.database_collection)
}

func (s *CartStore) CreateOne(ctx context.Context, cart cart_types.CartSchema) (*mongo.InsertOneResult, error) {
	col := s.getCartCollection()

	res, err := col.InsertOne(ctx, cart)
	if err != nil {
		return &mongo.InsertOneResult{}, err
	}

	return res, nil
}

func (s *CartStore) GetOne(ctx context.Context, id primitive.ObjectID) (cart_types.CartSchema, error) {
	col := s.getCartCollection()
	filter := bson.D{{Key: "_id", Value: id}}

	var res cart_types.CartSchema
	err := col.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return cart_types.CartSchema{}, err
	}

	return res, nil
}
