package brand

import (
	"context"
	"errors"

	"github.com/Kiyosh31/ms-ecommerce/product-service/cmd/internal/domain"
	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/cmd/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s Service) Create(ctx context.Context, brand domain.BrandSchema) (*productPb.Brand, error) {
	s.logger.Infof("create brand: %v", brand)

	res, err := s.repository.Create(ctx, brand)
	if err != nil {
		s.logger.Error(err)
		return &productPb.Brand{}, err
	}

	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return &productPb.Brand{}, errors.New("error getting new category id")
	}
	brand.ID = id

	s.logger.Infof("create brand finished with: %v", brand)
	return &productPb.Brand{
		Id:          brand.ID.Hex(),
		Name:        brand.Name,
		Description: brand.Description,
	}, nil
}
