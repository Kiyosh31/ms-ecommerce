package product

import (
	"context"

	"github.com/Kiyosh31/ms-ecommerce/product-service/cmd/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) Create(ctx context.Context, product domain.ProductSchema) (*mongo.InsertOneResult, error) {
	res, err := r.collection.InsertOne(ctx, product)
	if err != nil {
		return nil, err
	}

	return res, nil
}
