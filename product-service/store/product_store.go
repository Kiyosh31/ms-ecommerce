package store

import (
	"context"

	"github.com/Kiyosh31/ms-ecommerce/product-service/product_types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductStore struct {
	client              *mongo.Client
	database_name       string
	database_collection string
}

func NewProductStore(
	client *mongo.Client,
	database_name string,
	database_collection string,
) *ProductStore {
	return &ProductStore{
		client:              client,
		database_name:       database_name,
		database_collection: database_collection,
	}
}

func (s *ProductStore) getCollection() *mongo.Collection {
	return s.client.Database(s.database_name).Collection(s.database_collection)
}

func (s *ProductStore) CreateOne(ctx context.Context, user product_types.ProductSchema) (*mongo.InsertOneResult, error) {
	col := s.getCollection()

	res, err := col.InsertOne(ctx, user)
	if err != nil {
		return &mongo.InsertOneResult{}, err
	}

	return res, nil
}

func (s *ProductStore) GetOne(ctx context.Context, id primitive.ObjectID) (product_types.ProductSchema, error) {
	col := s.getCollection()
	filter := bson.D{{Key: "_id", Value: id}}

	var res product_types.ProductSchema
	err := col.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return product_types.ProductSchema{}, err
	}

	return res, nil
}

func (s *ProductStore) GetAll(ctx context.Context) ([]product_types.ProductSchema, error) {
	col := s.getCollection()

	// Find all documents
	cursor, err := col.Find(ctx, bson.D{})
	if err != nil {
		return []product_types.ProductSchema{}, err
	}
	defer cursor.Close(ctx)

	// Iterate over the cursor and decode results
	var results []product_types.ProductSchema
	for cursor.Next(ctx) {
		var result product_types.ProductSchema
		if err := cursor.Decode(&result); err != nil {
			return []product_types.ProductSchema{}, err
		}
		results = append(results, result)
	}
	if err := cursor.Err(); err != nil {
		return []product_types.ProductSchema{}, err
	}

	return results, nil
}

func (s *ProductStore) UpdateOne(ctx context.Context, userToUpdate product_types.ProductSchema) (*mongo.UpdateResult, error) {
	col := s.getCollection()
	filter := bson.D{{Key: "_id", Value: userToUpdate.ID}}
	update := bson.D{{Key: "$set", Value: userToUpdate}}

	res, err := col.UpdateOne(ctx, filter, update)
	if err != nil {
		return &mongo.UpdateResult{}, err
	}

	return res, nil
}

func (s *ProductStore) DeleteOne(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	col := s.getCollection()
	filter := bson.D{{Key: "_id", Value: id}}

	res, err := col.DeleteOne(ctx, filter)
	if err != nil {
		return &mongo.DeleteResult{}, err
	}

	return res, nil
}
