package store

import (
	"context"

	"github.com/Kiyosh31/ms-ecommerce/user-service/types"
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

func (s *UserStore) CreateOne(ctx context.Context, user types.UserSchema) (*mongo.InsertOneResult, error) {
	col := s.getUserCollection()

	res, err := col.InsertOne(ctx, user)
	if err != nil {
		return &mongo.InsertOneResult{}, err
	}

	return res, nil
}

func (s *UserStore) GetOne(ctx context.Context, id primitive.ObjectID) (types.UserSchema, error) {
	col := s.getUserCollection()
	filter := bson.D{{Key: "_id", Value: id}}

	var res types.UserSchema
	err := col.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return types.UserSchema{}, err
	}

	return res, nil
}

func (s *UserStore) GetOneByEmail(ctx context.Context, email string) (types.UserSchema, error) {
	col := s.getUserCollection()
	filter := bson.D{{Key: "email", Value: email}}

	var res types.UserSchema
	err := col.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return types.UserSchema{}, err
	}

	return res, nil
}

func (s *UserStore) UpdateOne(ctx context.Context, userToUpdate types.UserSchema) (*mongo.UpdateResult, error) {
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
	filter := bson.D{{Key: "_id", Value: id}}

	res, err := col.DeleteOne(ctx, filter)
	if err != nil {
		return &mongo.DeleteResult{}, err
	}

	return res, nil
}
