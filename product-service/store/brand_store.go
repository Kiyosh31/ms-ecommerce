package store

import (
	"context"
	"errors"

	"github.com/Kiyosh31/ms-ecommerce/product-service/product_types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BrandStore struct {
	client        *mongo.Client
	database_name string
	col           *mongo.Collection
}

func NewBrandStore(
	client *mongo.Client,
	database_name string,
	database_collection_name string,
) *BrandStore {
	return &BrandStore{
		client:        client,
		database_name: database_name,
		col:           client.Database(database_name).Collection(database_collection_name),
	}
}

func (s *BrandStore) CreateOne(ctx context.Context, user product_types.BrandSchema) (*mongo.InsertOneResult, error) {
	res, err := s.col.InsertOne(ctx, user)
	if err != nil {
		return &mongo.InsertOneResult{}, err
	}

	return res, nil
}

func (s *BrandStore) GetOne(ctx context.Context, id primitive.ObjectID) (product_types.BrandSchema, error) {
	filter := bson.D{{Key: "_id", Value: id}}

	var res product_types.BrandSchema
	err := s.col.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return product_types.BrandSchema{}, err
	}

	return res, nil
}

func (s *BrandStore) BrandExists(ctx context.Context, id primitive.ObjectID) (product_types.BrandSchema, bool, error) {
	filter := bson.D{{Key: "_id", Value: id}}

	var res product_types.BrandSchema
	err := s.col.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return res, false, nil // User not found
		}
		return res, false, err // Other error occurred
	}

	return res, true, nil // User found
}

func (s *BrandStore) GetAll(ctx context.Context) ([]product_types.BrandSchema, error) {

	// Find all documents
	cursor, err := s.col.Find(ctx, bson.D{})
	if err != nil {
		return []product_types.BrandSchema{}, err
	}
	defer cursor.Close(ctx)

	// Iterate over the cursor and decode results
	var results []product_types.BrandSchema
	for cursor.Next(ctx) {
		var result product_types.BrandSchema
		if err := cursor.Decode(&result); err != nil {
			return []product_types.BrandSchema{}, err
		}
		results = append(results, result)
	}
	if err := cursor.Err(); err != nil {
		return []product_types.BrandSchema{}, err
	}

	return results, nil
}

func (s *BrandStore) UpdateOne(ctx context.Context, userToUpdate product_types.BrandSchema) (*mongo.UpdateResult, error) {
	filter := bson.D{{Key: "_id", Value: userToUpdate.ID}}
	update := bson.D{{Key: "$set", Value: userToUpdate}}

	res, err := s.col.UpdateOne(ctx, filter, update)
	if err != nil {
		return &mongo.UpdateResult{}, err
	}

	return res, nil
}

func (s *BrandStore) DeleteOne(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.D{{Key: "_id", Value: id}}

	res, err := s.col.DeleteOne(ctx, filter)
	if err != nil {
		return &mongo.DeleteResult{}, err
	}

	return res, nil
}
