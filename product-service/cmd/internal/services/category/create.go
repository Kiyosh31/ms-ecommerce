package category

import (
	"context"
	"errors"

	"github.com/Kiyosh31/ms-ecommerce/product-service/cmd/internal/domain"
	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/cmd/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s Service) Create(ctx context.Context, category domain.CategorySchema) (*productPb.Category, error) {
	s.logger.Infof("create category: %v", category)

	res, err := s.repository.Create(ctx, category)
	if err != nil {
		s.logger.Error(err)
		return &productPb.Category{}, err
	}

	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return &productPb.Category{}, errors.New("error getting new category id")
	}
	category.ID = id

	s.logger.Infof("create category finished with: %v", category)
	return &productPb.Category{
		Id:          category.ID.Hex(),
		Name:        category.Name,
		Description: category.Description,
	}, nil
}
