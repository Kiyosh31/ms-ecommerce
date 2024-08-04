package store

import (
	"context"
	"errors"

	ordertypes "github.com/Kiyosh31/ms-ecommerce/order-service/order_types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderStore struct {
	client        *mongo.Client
	database_name string
	col           *mongo.Collection
}

func NewOrderStore(
	client *mongo.Client,
	database_name string,
	database_collection_name string,
) *OrderStore {
	return &OrderStore{
		client:        client,
		database_name: database_name,
		col:           client.Database(database_name).Collection(database_collection_name),
	}
}

func (s *OrderStore) CreateOne(ctx context.Context, user ordertypes.OrderSchema) (*mongo.InsertOneResult, error) {
	res, err := s.col.InsertOne(ctx, user)
	if err != nil {
		return &mongo.InsertOneResult{}, err
	}

	return res, nil
}

func (s *OrderStore) GetOne(ctx context.Context, id primitive.ObjectID) (ordertypes.OrderSchema, error) {
	filter := bson.D{{Key: "_id", Value: id}}

	var res ordertypes.OrderSchema
	err := s.col.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return ordertypes.OrderSchema{}, err
	}

	return res, nil
}

func (s *OrderStore) ProductExists(ctx context.Context, id primitive.ObjectID) (ordertypes.OrderSchema, bool, error) {
	filter := bson.D{{Key: "_id", Value: id}}

	var res ordertypes.OrderSchema
	err := s.col.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return res, false, nil // User not found
		}
		return res, false, err // Other error occurred
	}

	return res, true, nil // User found
}

func (s *OrderStore) GetAll(ctx context.Context) ([]ordertypes.OrderSchema, error) {

	// Find all documents
	cursor, err := s.col.Find(ctx, bson.D{})
	if err != nil {
		return []ordertypes.OrderSchema{}, err
	}
	defer cursor.Close(ctx)

	// Iterate over the cursor and decode results
	var results []ordertypes.OrderSchema
	for cursor.Next(ctx) {
		var result ordertypes.OrderSchema
		if err := cursor.Decode(&result); err != nil {
			return []ordertypes.OrderSchema{}, err
		}
		results = append(results, result)
	}
	if err := cursor.Err(); err != nil {
		return []ordertypes.OrderSchema{}, err
	}

	return results, nil
}

func (s *OrderStore) UpdateOne(ctx context.Context, userToUpdate ordertypes.OrderSchema) (*mongo.UpdateResult, error) {
	filter := bson.D{{Key: "_id", Value: userToUpdate.ID}}
	update := bson.D{{Key: "$set", Value: userToUpdate}}

	res, err := s.col.UpdateOne(ctx, filter, update)
	if err != nil {
		return &mongo.UpdateResult{}, err
	}

	return res, nil
}

func (s *OrderStore) DeleteOne(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
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
