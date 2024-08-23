package user

import (
	"context"

	"github.com/Kiyosh31/ms-ecommerce/user-service/cmd/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) Create(ctx context.Context, user domain.UserSchema) (*mongo.InsertOneResult, error) {
	res, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return res, nil
}
