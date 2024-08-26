package payment

import "go.mongodb.org/mongo-driver/mongo"

type Repository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewPaymentRepository(
	client *mongo.Client,
	databaseName string,
	collectionName string,
) *Repository {
	return &Repository{
		client:     client,
		collection: client.Database(databaseName).Collection(collectionName),
	}
}
