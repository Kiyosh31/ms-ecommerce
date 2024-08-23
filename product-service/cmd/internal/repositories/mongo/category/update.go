package category

import (
	"context"

	"github.com/Kiyosh31/ms-ecommerce/product-service/cmd/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) Update(ctx context.Context, category domain.CategorySchema) (mongo.UpdateResult, error) {
	filter := bson.D{{Key: "_id", Value: category.ID}}
	update := bson.D{{Key: "$set", Value: category}}

	res, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return mongo.UpdateResult{}, err
	}

	return *res, nil
}
