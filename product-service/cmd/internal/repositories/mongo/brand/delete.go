package brand

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) Delete(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.D{{Key: "_id", Value: id}}

	res, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return &mongo.DeleteResult{}, err
	}

	// Check if a document was actually deleted
	if res.DeletedCount == 0 {
		// Handle the case where no document was found
		return &mongo.DeleteResult{}, errors.New("document not found")
	}

	return res, nil
}
