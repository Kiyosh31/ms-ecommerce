package store

import (
	"context"
	"errors"

	"github.com/Kiyosh31/ms-ecommerce/product-service/product_types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductStore struct {
	client        *mongo.Client
	database_name string
	col           *mongo.Collection
}

func NewProductStore(
	client *mongo.Client,
	database_name string,
	database_collection_name string,
) *ProductStore {
	return &ProductStore{
		client:        client,
		database_name: database_name,
		col:           client.Database(database_name).Collection(database_collection_name),
	}
}

func (s *ProductStore) CreateOne(ctx context.Context, user product_types.ProductSchema) (*mongo.InsertOneResult, error) {
	res, err := s.col.InsertOne(ctx, user)
	if err != nil {
		return &mongo.InsertOneResult{}, err
	}

	return res, nil
}

func (s *ProductStore) GetOne(ctx context.Context, id primitive.ObjectID) (product_types.ProductSchema, error) {
	filter := bson.D{{Key: "_id", Value: id}}

	var res product_types.ProductSchema
	err := s.col.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return product_types.ProductSchema{}, err
	}

	return res, nil
}

func (s *ProductStore) ProductExists(ctx context.Context, id primitive.ObjectID) (product_types.ProductSchema, bool, error) {
	filter := bson.D{{Key: "_id", Value: id}}

	var res product_types.ProductSchema
	err := s.col.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return res, false, nil // User not found
		}
		return res, false, err // Other error occurred
	}

	return res, true, nil // User found
}

func (s *ProductStore) GetAll(ctx context.Context) ([]product_types.ProductSchema, error) {

	// Find all documents
	cursor, err := s.col.Find(ctx, bson.D{})
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
	filter := bson.D{{Key: "_id", Value: userToUpdate.ID}}
	update := bson.D{{Key: "$set", Value: userToUpdate}}

	res, err := s.col.UpdateOne(ctx, filter, update)
	if err != nil {
		return &mongo.UpdateResult{}, err
	}

	return res, nil
}

func (s *ProductStore) DeleteOne(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.D{{Key: "_id", Value: id}}

	res, err := s.col.DeleteOne(ctx, filter)
	if err != nil {
		return &mongo.DeleteResult{}, err
	}

	// Check if a document was actually deleted
	if res.DeletedCount == 0 {
		// Handle the case where no document was found
		return nil, errors.New("document not found")
	}

	return res, nil
}
