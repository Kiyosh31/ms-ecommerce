package product

import (
	"context"

	"github.com/Kiyosh31/ms-ecommerce/product-service/cmd/internal/domain"
	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/cmd/proto"
)

func (s Service) Create(ctx context.Context, product domain.ProductSchema) (*productPb.Product, error) {
	return &productPb.Product{}, nil
	// s.logger.Infof("Create product incoming request: %v", product)

	// categoryId, err := database.GetMongoId(product.Category.ID.Hex())
	// if err != nil {
	// 	s.logger.Errorf("error getting categoryId: %v", err)
	// 	return &productPb.Product{}, err
	// }

	// brandId, err := database.GetMongoId(product.Brand.ID.Hex())
	// if err != nil {
	// 	s.logger.Errorf("error getting brandId: %v", err)
	// 	return &productPb.Product{}, err
	// }

	// // check category exists
	// _, exists, err := s.CategoryStore.CategoryExists(ctx, categoryId)
	// if err != nil {
	// 	s.logger.Errorf("error finding existing category: %v", err)
	// 	return &productPb.Product{}, err
	// }
	// if !exists {
	// 	s.logger.Errorf("error category dont exists: %v", err)
	// 	return &productPb.Product{}, errors.New("category dont exist")
	// }

	// // check brand exists
	// _, exists, err = s.BrandStore.BrandExists(ctx, brandId)
	// if err != nil {
	// 	s.logger.Errorf("error finding existing brand: %v", err)
	// 	return &productPb.Product{}, err
	// }
	// if !exists {
	// 	s.logger.Errorf("error category dont exists: %v", err)
	// 	return &productPb.Product{}, errors.New("brand dont exist")
	// }

	// // create product schema for db
	// productDto, err := createProductSchemaDto(in.GetProduct())
	// if err != nil {
	// 	s.logger.Errorf("error creating schema: %v", err)
	// 	return &productPb.Product{}, err
	// }

	// // saving product to db
	// createdProduct, err := s.repository.Create(ctx, productDto)
	// if err != nil {
	// 	s.logger.Errorf("error creating product: %v", err)
	// 	return &productPb.Product{}, err
	// }

	// // getting productID created in db
	// id, ok := createdProduct.InsertedID.(primitive.ObjectID)
	// if !ok {
	// 	s.logger.Errorf("error getting id: %v", err)
	// 	return &productPb.Product{}, fmt.Errorf("failed to parse _id to string")
	// }
	// product.ID = id.Hex()

	// // return response
	// s.logger.Infof("create product request finished: %v", createdProduct)
	// return &productPb.Product{}, nil
}
