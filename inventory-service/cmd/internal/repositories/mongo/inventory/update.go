package inventory

import (
	"context"

	"github.com/Kiyosh31/ms-ecommerce/inventory-service/cmd/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) Update(ctx context.Context, user domain.InventorySchema) (mongo.UpdateResult, error) {
	filter := bson.D{{Key: "_id", Value: user.ID}}
	update := bson.D{{Key: "$set", Value: user}}

	res, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return mongo.UpdateResult{}, err
	}

	return *res, nil
}
