package user

import (
	"context"
	"errors"

	"github.com/Kiyosh31/ms-ecommerce/user-service/cmd/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r Repository) GetByEmail(ctx context.Context, email string) (domain.UserSchema, bool, error) {
	filter := bson.D{{Key: "email", Value: email}}

	var res domain.UserSchema
	err := r.collection.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.UserSchema{}, false, nil // User not found
		}
		return domain.UserSchema{}, false, err // Other error occurred
	}

	return res, true, nil
}
