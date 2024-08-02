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
	client              *mongo.Client
	database_name       string
	database_collection string
}

func NewUserStore(
	client *mongo.Client,
	database_name string,
	database_collection string,
) *UserStore {
	return &UserStore{
		client:              client,
		database_name:       database_name,
		database_collection: database_collection,
	}
}

func (s *UserStore) getUserCollection() *mongo.Collection {
	return s.client.Database(s.database_name).Collection(s.database_collection)
}

func (s *UserStore) CreateOne(ctx context.Context, user user_types.UserSchema) (*mongo.InsertOneResult, error) {
	col := s.getUserCollection()

	res, err := col.InsertOne(ctx, user)
	if err != nil {
		return &mongo.InsertOneResult{}, err
	}

	return res, nil
}

func (s *UserStore) GetOne(ctx context.Context, id primitive.ObjectID) (user_types.UserSchema, error) {
	col := s.getUserCollection()
	filter := bson.D{
		{Key: "_id", Value: id},
		{Key: "isActive", Value: true},
	}

	var res user_types.UserSchema
	err := col.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return user_types.UserSchema{}, err
	}

	return res, nil
}

func (s *UserStore) GetOneDeactivated(ctx context.Context, email string) (user_types.UserSchema, error) {
	col := s.getUserCollection()
	filter := bson.D{
		{Key: "email", Value: email},
		{Key: "isActive", Value: false},
	}

	var res user_types.UserSchema
	err := col.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return user_types.UserSchema{}, err
	}

	return res, nil
}

func (s *UserStore) GetOneByEmail(ctx context.Context, email string) (user_types.UserSchema, error) {
	col := s.getUserCollection()
	filter := bson.D{
		{Key: "email", Value: email},
		{Key: "isActive", Value: true},
	}

	var res user_types.UserSchema
	err := col.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return user_types.UserSchema{}, err
	}

	return res, nil
}

func (s *UserStore) UserExists(ctx context.Context, email string) (bool, error) {
	col := s.getUserCollection()
	filter := bson.D{
		{Key: "email", Value: email},
		{Key: "isActive", Value: true},
	}

	var res user_types.UserSchema
	err := col.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil // User not found
		}
		return false, err // Other error occurred
	}

	return true, nil // User found
}

func (s *UserStore) UpdateOne(ctx context.Context, userToUpdate user_types.UserSchema) (*mongo.UpdateResult, error) {
	col := s.getUserCollection()
	filter := bson.D{{Key: "_id", Value: userToUpdate.ID}}
	update := bson.D{{Key: "$set", Value: userToUpdate}}

	res, err := col.UpdateOne(ctx, filter, update)
	if err != nil {
		return &mongo.UpdateResult{}, err
	}

	return res, nil
}

func (s *UserStore) DeleteOne(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	col := s.getUserCollection()
	filter := bson.D{
		{Key: "_id", Value: id},
		{Key: "isActive", Value: true},
	}

	res, err := col.DeleteOne(ctx, filter)
	if err != nil {
		return &mongo.DeleteResult{}, err
	}

	return res, nil
}
