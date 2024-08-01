package store

import (
	"context"

	stocktypes "github.com/Kiyosh31/ms-ecommerce/stock-service/stock_types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type StockStore struct {
	client              *mongo.Client
	database_name       string
	database_collection string
}

func NewStockStore(client *mongo.Client, databaseName string, databaseCollection string) *StockStore {
	return &StockStore{
		client:              client,
		database_name:       databaseName,
		database_collection: databaseCollection,
	}
}

func (s *StockStore) getCollection() *mongo.Collection {
	return s.client.Database(s.database_name).Collection(s.database_collection)
}

func (s *StockStore) CreateOne(ctx context.Context, stock stocktypes.StockSchema) (*mongo.InsertOneResult, error) {
	col := s.getCollection()

	res, err := col.InsertOne(ctx, stock)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return &mongo.InsertOneResult{}, err
		}
		return &mongo.InsertOneResult{}, err
	}

	return res, nil
}

func (s *StockStore) GetOne(ctx context.Context, id primitive.ObjectID) (stocktypes.StockSchema, error) {
	col := s.getCollection()
	filter := bson.D{{Key: "_id", Value: id}}

	var res stocktypes.StockSchema
	err := col.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return stocktypes.StockSchema{}, err
		}
		return stocktypes.StockSchema{}, err
	}

	return res, nil
}

func (s *StockStore) UpdateOne(ctx context.Context, stockToUpdate stocktypes.StockSchema) (*mongo.UpdateResult, error) {
	col := s.getCollection()
	filter := bson.D{{Key: "_id", Value: stockToUpdate.ID}}
	update := bson.D{{Key: "$set", Value: stockToUpdate}}

	res, err := col.UpdateOne(ctx, filter, update)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return &mongo.UpdateResult{}, err
		}
		return &mongo.UpdateResult{}, err
	}

	return res, nil
}

func (s *StockStore) DeleteOne(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	col := s.getCollection()
	filter := bson.D{{Key: "_id", Value: id}}

	res, err := col.DeleteOne(ctx, filter)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return &mongo.DeleteResult{}, err
		}
		return &mongo.DeleteResult{}, err
	}

	return res, nil
}
