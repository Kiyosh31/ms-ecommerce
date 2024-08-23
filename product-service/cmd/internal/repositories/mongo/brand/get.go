package brand

import (
	"context"
	"errors"

	"github.com/Kiyosh31/ms-ecommerce/product-service/cmd/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) Get(ctx context.Context, id primitive.ObjectID) (domain.BrandSchema, bool, error) {
	filter := bson.D{{Key: "_id", Value: id}}

	var res domain.BrandSchema
	err := r.collection.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.BrandSchema{}, false, nil // not found
		}
		return domain.BrandSchema{}, false, err // Other error occurred
	}

	return res, true, nil
}
