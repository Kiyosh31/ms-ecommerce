package user

import "go.mongodb.org/mongo-driver/mongo"

type Repository struct {
	client        *mongo.Client
	database_name string
	collection    *mongo.Collection
}

func NewUserRepository(
	client *mongo.Client,
	database_name string,
	collection_name string,
) *Repository {
	return &Repository{
		client:        client,
		database_name: database_name,
		collection:    client.Database(database_name).Collection(collection_name),
	}
}
