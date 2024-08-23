package product

import (
	"context"
	"errors"

	"github.com/Kiyosh31/ms-ecommerce/product-service/cmd/internal/domain"
	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/cmd/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s Service) Create(ctx context.Context, product domain.ProductSchema) (*productPb.Product, error) {
	s.logger.Infof("create product: %v", product)

	res, err := s.repository.Create(ctx, product)
	if err != nil {
		s.logger.Error(err)
		return &productPb.Product{}, err
	}

	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return &productPb.Product{}, errors.New("error getting new product id")
	}
	product.ID = id

	s.logger.Infof("create product finished with: %v", product)
	return &productPb.Product{
		Id:          product.ID.Hex(),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Category:    &productPb.Category{},
		Brand:       &productPb.Brand{},
		Images:      product.Images,
		Attributes:  &productPb.Attributes{},
		Inventory:   &productPb.Inventory{},
	}, nil
}
