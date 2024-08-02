package store

import (
	"context"
	"errors"

	"github.com/Kiyosh31/ms-ecommerce/product-service/product_types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryStore struct {
	client        *mongo.Client
	database_name string
	col           *mongo.Collection
}

func NewCategoryStore(
	client *mongo.Client,
	database_name string,
	database_collection_name string,
) *CategoryStore {
	return &CategoryStore{
		client:        client,
		database_name: database_name,
		col:           client.Database(database_name).Collection(database_collection_name),
	}
}

func (s *CategoryStore) CreateOne(ctx context.Context, user product_types.CategorySchema) (*mongo.InsertOneResult, error) {
	res, err := s.col.InsertOne(ctx, user)
	if err != nil {
		return &mongo.InsertOneResult{}, err
	}

	return res, nil
}

func (s *CategoryStore) GetOne(ctx context.Context, id primitive.ObjectID) (product_types.CategorySchema, error) {
	filter := bson.D{{Key: "_id", Value: id}}

	var res product_types.CategorySchema
	err := s.col.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return product_types.CategorySchema{}, err
	}

	return res, nil
}

func (s *CategoryStore) CategoryExists(ctx context.Context, id primitive.ObjectID) (product_types.CategorySchema, bool, error) {
	filter := bson.D{{Key: "_id", Value: id}}

	var res product_types.CategorySchema
	err := s.col.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return res, false, nil // User not found
		}
		return res, false, err // Other error occurred
	}

	return res, true, nil // User found
}

func (s *CategoryStore) GetAll(ctx context.Context) ([]product_types.CategorySchema, error) {

	// Find all documents
	cursor, err := s.col.Find(ctx, bson.D{})
	if err != nil {
		return []product_types.CategorySchema{}, err
	}
	defer cursor.Close(ctx)

	// Iterate over the cursor and decode results
	var results []product_types.CategorySchema
	for cursor.Next(ctx) {
		var result product_types.CategorySchema
		if err := cursor.Decode(&result); err != nil {
			return []product_types.CategorySchema{}, err
		}
		results = append(results, result)
	}
	if err := cursor.Err(); err != nil {
		return []product_types.CategorySchema{}, err
	}

	return results, nil
}

func (s *CategoryStore) UpdateOne(ctx context.Context, userToUpdate product_types.CategorySchema) (*mongo.UpdateResult, error) {
	filter := bson.D{{Key: "_id", Value: userToUpdate.ID}}
	update := bson.D{{Key: "$set", Value: userToUpdate}}

	res, err := s.col.UpdateOne(ctx, filter, update)
	if err != nil {
		return &mongo.UpdateResult{}, err
	}

	return res, nil
}

func (s *CategoryStore) DeleteOne(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.D{{Key: "_id", Value: id}}

	res, err := s.col.DeleteOne(ctx, filter)
	if err != nil {
		return &mongo.DeleteResult{}, err
	}

	return res, nil
}
