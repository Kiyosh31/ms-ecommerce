package category

import (
	"context"

	"github.com/Kiyosh31/ms-ecommerce/product-service/cmd/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) Create(ctx context.Context, category domain.CategorySchema) (*mongo.InsertOneResult, error) {
	res, err := r.collection.InsertOne(ctx, category)
	if err != nil {
		return nil, err
	}

	return res, nil
}
