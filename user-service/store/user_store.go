package store

import (
	"context"
	"errors"

	"github.com/Kiyosh31/ms-ecommerce/user-service/user_types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStore struct {
	client        *mongo.Client
	database_name string
	col           *mongo.Collection
}

func NewUserStore(
	client *mongo.Client,
	database_name string,
	database_collection_name string,
) *UserStore {
	return &UserStore{
		client:        client,
		database_name: database_name,
		col:           client.Database(database_name).Collection(database_collection_name),
	}
}

func (s *UserStore) CreateOne(ctx context.Context, user user_types.UserSchema) (*mongo.InsertOneResult, error) {
	res, err := s.col.InsertOne(ctx, user)
	if err != nil {
		return &mongo.InsertOneResult{}, err
	}

	return res, nil
}

func (s *UserStore) GetOne(ctx context.Context, id primitive.ObjectID) (user_types.UserSchema, error) {
	filter := bson.D{
		{Key: "_id", Value: id},
		{Key: "isActive", Value: true},
	}

	var res user_types.UserSchema
	err := s.col.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return user_types.UserSchema{}, err
	}

	return res, nil
}

func (s *UserStore) GetOneDeactivated(ctx context.Context, email string) (user_types.UserSchema, error) {
	filter := bson.D{
		{Key: "email", Value: email},
		{Key: "isActive", Value: false},
	}

	var res user_types.UserSchema
	err := s.col.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return user_types.UserSchema{}, err
	}

	return res, nil
}

func (s *UserStore) GetOneByEmail(ctx context.Context, email string) (user_types.UserSchema, error) {
	filter := bson.D{
		{Key: "email", Value: email},
		{Key: "isActive", Value: true},
	}

	var res user_types.UserSchema
	err := s.col.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return user_types.UserSchema{}, err
	}

	return res, nil
}

func (s *UserStore) UserExists(ctx context.Context, email string) (bool, error) {
	filter := bson.D{
		{Key: "email", Value: email},
		{Key: "isActive", Value: true},
	}

	var res user_types.UserSchema
	err := s.col.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil // User not found
		}
		return false, err // Other error occurred
	}

	return true, nil // User found
}

func (s *UserStore) UpdateOne(ctx context.Context, userToUpdate user_types.UserSchema) (*mongo.UpdateResult, error) {
	filter := bson.D{{Key: "_id", Value: userToUpdate.ID}}
	update := bson.D{{Key: "$set", Value: userToUpdate}}

	res, err := s.col.UpdateOne(ctx, filter, update)
	if err != nil {
		return &mongo.UpdateResult{}, err
	}

	return res, nil
}

func (s *UserStore) DeleteOne(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.D{
		{Key: "_id", Value: id},
		{Key: "isActive", Value: true},
	}

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
