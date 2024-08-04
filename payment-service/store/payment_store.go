package store

import (
	"context"
	"errors"

	"github.com/Kiyosh31/ms-ecommerce/payment-service/payment_types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PaymentStore struct {
	client        *mongo.Client
	database_name string
	col           *mongo.Collection
}

func NewPaymentStore(
	client *mongo.Client,
	database_name string,
	database_collection_name string,
) *PaymentStore {
	return &PaymentStore{
		client:        client,
		database_name: database_name,
		col:           client.Database(database_name).Collection(database_collection_name),
	}
}

func (s *PaymentStore) CreateOne(ctx context.Context, payment payment_types.PaymentSchema) (*mongo.InsertOneResult, error) {
	res, err := s.col.InsertOne(ctx, payment)
	if err != nil {
		return &mongo.InsertOneResult{}, err
	}

	return res, nil
}

func (s *PaymentStore) GetOne(ctx context.Context, id primitive.ObjectID) (payment_types.PaymentSchema, error) {
	filter := bson.D{{Key: "_id", Value: id}}

	var res payment_types.PaymentSchema
	err := s.col.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return payment_types.PaymentSchema{}, err
	}

	return res, nil
}

func (s *PaymentStore) GetOneDeactivated(ctx context.Context, email string) (payment_types.PaymentSchema, error) {
	filter := bson.D{
		{Key: "email", Value: email},
		{Key: "isActive", Value: false},
	}

	var res payment_types.PaymentSchema
	err := s.col.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return payment_types.PaymentSchema{}, err
	}

	return res, nil
}

func (s *PaymentStore) GetOneByEmail(ctx context.Context, email string) (payment_types.PaymentSchema, error) {
	filter := bson.D{{Key: "email", Value: email}}

	var res payment_types.PaymentSchema
	err := s.col.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return payment_types.PaymentSchema{}, err
	}

	return res, nil
}

func (s *PaymentStore) PaymentExists(ctx context.Context, email string) (bool, error) {
	filter := bson.D{{Key: "email", Value: email}}

	var res payment_types.PaymentSchema
	err := s.col.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil // not found
		}
		return false, err // Other error occurred
	}

	return true, nil // found
}

func (s *PaymentStore) UpdateOne(ctx context.Context, paymentToUpdate payment_types.PaymentSchema) (*mongo.UpdateResult, error) {
	filter := bson.D{{Key: "_id", Value: paymentToUpdate.ID}}
	update := bson.D{{Key: "$set", Value: paymentToUpdate}}

	res, err := s.col.UpdateOne(ctx, filter, update)
	if err != nil {
		return &mongo.UpdateResult{}, err
	}

	return res, nil
}

func (s *PaymentStore) DeleteOne(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
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
