package ports

import (
	"context"

	"github.com/Kiyosh31/ms-ecommerce/user-service/cmd/internal/domain"
	userPb "github.com/Kiyosh31/ms-ecommerce/user-service/cmd/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService interface {
	Create(ctx context.Context, new_user domain.UserSchema) (*userPb.User, error)
	Get(ctx context.Context, id string) (*userPb.User, error)
	Update(ctx context.Context, user_to_update domain.UserSchema) (*userPb.User, error)
	Deactivate(ctx context.Context, id string) (bool, error)
	Reactivate(ctx context.Context, email, password string) (*userPb.User, error)
}

type UserRepository interface {
	Create(ctx context.Context, user domain.UserSchema) (*mongo.InsertOneResult, error)
	Get(ctx context.Context, id primitive.ObjectID) (domain.UserSchema, bool, error)
	GetByEmail(ctx context.Context, email string) (domain.UserSchema, bool, error)
	Update(ctx context.Context, user domain.UserSchema) (mongo.UpdateResult, error)
}
