package category

import (
	"context"
	"errors"

	"github.com/Kiyosh31/ms-ecommerce/product-service/cmd/internal/domain"
	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/cmd/proto"
)

func (s Service) Update(ctx context.Context, categoryToUpdate domain.CategorySchema) (*productPb.Category, error) {
	s.logger.Infof("update category: %v", categoryToUpdate)

	// check if category exists
	_, exist, err := s.repository.Get(ctx, categoryToUpdate.ID)
	if err != nil {
		return &productPb.Category{}, err
	}
	if !exist {
		return &productPb.Category{}, errors.New("category not found")
	}

	_, err = s.repository.Update(ctx, categoryToUpdate)
	if err != nil {
		s.logger.Error(err)
		return &productPb.Category{}, err
	}

	s.logger.Infof("update category finished with: %v", categoryToUpdate)
	return &productPb.Category{
		Id:          categoryToUpdate.ID.Hex(),
		Name:        categoryToUpdate.Name,
		Description: categoryToUpdate.Description,
	}, nil
}
