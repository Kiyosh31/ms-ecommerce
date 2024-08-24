package brand

import (
	"context"
	"errors"

	"github.com/Kiyosh31/ms-ecommerce-common/database"
	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/cmd/proto"
)

func (s Service) Get(ctx context.Context, id string) (*productPb.Brand, error) {
	s.logger.Infof("get brand: %v", id)

	categoryId, err := database.GetMongoId(id)
	if err != nil {
		return &productPb.Brand{}, err
	}

	res, exist, err := s.repository.Get(ctx, categoryId)
	if err != nil {
		s.logger.Error(err)
		return &productPb.Brand{}, err
	}
	if !exist {
		return &productPb.Brand{}, errors.New("no brand found")
	}

	s.logger.Infof("get brand finished with: %v", res)
	return &productPb.Brand{
		Id:          res.ID.Hex(),
		Name:        res.Name,
		Description: res.Description,
	}, nil
}
