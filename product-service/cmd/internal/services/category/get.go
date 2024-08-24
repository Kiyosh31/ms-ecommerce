package category

import (
	"context"
	"errors"

	"github.com/Kiyosh31/ms-ecommerce-common/database"
	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/cmd/proto"
)

func (s Service) Get(ctx context.Context, id string) (*productPb.Category, error) {
	s.logger.Infof("get category: %v", id)

	categoryId, err := database.GetMongoId(id)
	if err != nil {
		return &productPb.Category{}, err
	}

	res, exist, err := s.repository.Get(ctx, categoryId)
	if err != nil {
		s.logger.Error(err)
		return &productPb.Category{}, err
	}
	if !exist {
		return &productPb.Category{}, errors.New("No category found")
	}

	s.logger.Infof("get category finished with: %v", res)
	return &productPb.Category{
		Id:          res.ID.Hex(),
		Name:        res.Name,
		Description: res.Description,
	}, nil
}
