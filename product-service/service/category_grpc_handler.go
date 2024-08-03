package service

import (
	"context"
	"fmt"

	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *ProductService) CreateCategory(ctx context.Context, in *productPb.CategoryRequest) (*productPb.CategoryResponse, error) {
	s.logger.Info("create category incoming request: %v", in)

	_, exists, err := s.CategoryStore.GetOneByName(ctx, in.GetCategory().GetName())
	if err != nil {
		s.logger.Errorf("error finding existing category: %v", err)
		return &productPb.CategoryResponse{}, err
	}
	if exists {
		s.logger.Errorf("error category already exists: %v", err)
		return &productPb.CategoryResponse{}, err
	}

	// create schema for DB
	categoryDto, err := createCategorySchemaDto(in.GetCategory())
	if err != nil {
		s.logger.Errorf("error creating category schema: %v", err)
		return &productPb.CategoryResponse{}, nil
	}

	// saving category to db
	createdCategory, err := s.CategoryStore.CreateOne(ctx, categoryDto)
	if err != nil {
		s.logger.Errorf("error creating category in db: %v", err)
		return &productPb.CategoryResponse{}, nil
	}

	// getting categoryID created in db
	id, ok := createdCategory.InsertedID.(primitive.ObjectID)
	if !ok {
		s.logger.Errorf("error getting id: %v", err)
		return &productPb.CategoryResponse{}, fmt.Errorf("failed to parse _id to string")
	}
	in.GetCategory().Id = id.Hex()

	// return response
	s.logger.Info("create category request finished: %v", createdCategory)
	return &productPb.CategoryResponse{
		Message:  "Category created successfully",
		Category: []*productPb.Category{in.GetCategory()},
	}, nil
}

func (s *ProductService) GetCategory(ctx context.Context, in *productPb.CategoryRequest) (*productPb.CategoryResponse, error) {
	return &productPb.CategoryResponse{}, nil

}

func (s *ProductService) UpdateCategory(ctx context.Context, in *productPb.CategoryRequest) (*productPb.CategoryResponse, error) {
	return &productPb.CategoryResponse{}, nil

}

func (s *ProductService) DeleteCategory(ctx context.Context, in *productPb.CategoryRequest) (*productPb.CategoryResponse, error) {
	return &productPb.CategoryResponse{}, nil

}
