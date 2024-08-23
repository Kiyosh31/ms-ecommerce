package ports

import (
	"context"

	"github.com/Kiyosh31/ms-ecommerce/product-service/cmd/internal/domain"
	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/cmd/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductService interface {
	Create(ctx context.Context, product domain.ProductSchema) (*productPb.Product, error)
	Get(ctx context.Context, id string) (*productPb.Product, error)
	Update(ctx context.Context, product domain.ProductSchema) (*productPb.Product, error)
	Delete(ctx context.Context, id string) (bool, error)
}

type ProductRepository interface {
	Create(ctx context.Context, product domain.ProductSchema) (*mongo.InsertOneResult, error)
	Get(ctx context.Context, id primitive.ObjectID) (domain.ProductSchema, bool, error)
	Update(ctx context.Context, product domain.ProductSchema) (mongo.UpdateResult, error)
	Delete(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error)
}

type CategoryService interface {
	Create(ctx context.Context, category domain.CategorySchema) (*productPb.Category, error)
	Get(ctx context.Context, id string) (*productPb.Category, error)
	Update(ctx context.Context, category domain.CategorySchema) (*productPb.Category, error)
	Delete(ctx context.Context, id string) (bool, error)
}

type CategoryRepository interface {
	Create(ctx context.Context, category domain.CategorySchema) (*mongo.InsertOneResult, error)
	Get(ctx context.Context, id primitive.ObjectID) (domain.CategorySchema, bool, error)
	Update(ctx context.Context, category domain.CategorySchema) (mongo.UpdateResult, error)
	Delete(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error)
}

type BrandService interface {
	Create(ctx context.Context, brand domain.BrandSchema) (*productPb.Brand, error)
	Get(ctx context.Context, id string) (*productPb.Brand, error)
	Update(ctx context.Context, brand domain.BrandSchema) (*productPb.Brand, error)
	Delete(ctx context.Context, id string) (bool, error)
}

type BrandRepository interface {
	Create(ctx context.Context, product domain.BrandSchema) (*mongo.InsertOneResult, error)
	Get(ctx context.Context, id primitive.ObjectID) (domain.BrandSchema, bool, error)
	Update(ctx context.Context, product domain.BrandSchema) (mongo.UpdateResult, error)
	Delete(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error)
}
