package store

import (
	"context"
	"errors"

	"github.com/Kiyosh31/ms-ecommerce/inventory-service/inventory_types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type InventoryStore struct {
	client        *mongo.Client
	database_name string
	col           *mongo.Collection
}

func NewInventoryStore(
	client *mongo.Client,
	database_name string,
	database_collection_name string,
) *InventoryStore {
	return &InventoryStore{
		client:        client,
		database_name: database_name,
		col:           client.Database(database_name).Collection(database_collection_name),
	}
}

func (s *InventoryStore) CreateOne(ctx context.Context, user inventory_types.InventorySchema) (*mongo.InsertOneResult, error) {
	res, err := s.col.InsertOne(ctx, user)
	if err != nil {
		return &mongo.InsertOneResult{}, err
	}

	return res, nil
}

func (s *InventoryStore) GetOne(ctx context.Context, id primitive.ObjectID) (inventory_types.InventorySchema, error) {
	filter := bson.D{{Key: "_id", Value: id}}

	var res inventory_types.InventorySchema
	err := s.col.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return inventory_types.InventorySchema{}, err
	}

	return res, nil
}

func (s *InventoryStore) InventoryExists(ctx context.Context, id primitive.ObjectID) (inventory_types.InventorySchema, bool, error) {
	filter := bson.D{{Key: "_id", Value: id}}

	var res inventory_types.InventorySchema
	err := s.col.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return res, false, nil // User not found
		}
		return res, false, err // Other error occurred
	}

	return res, true, nil // User found
}

func (s *InventoryStore) UpdateOne(ctx context.Context, userToUpdate inventory_types.InventorySchema) (*mongo.UpdateResult, error) {
	filter := bson.D{{Key: "_id", Value: userToUpdate.ID}}
	update := bson.D{{Key: "$set", Value: userToUpdate}}

	res, err := s.col.UpdateOne(ctx, filter, update)
	if err != nil {
		return &mongo.UpdateResult{}, err
	}

	return res, nil
}

func (s *InventoryStore) DeleteOne(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
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
