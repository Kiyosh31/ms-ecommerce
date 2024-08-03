package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/Kiyosh31/ms-ecommerce-common/database"
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
		return &productPb.CategoryResponse{}, errors.New("category already exists")
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
	s.logger.Info("get category incoming request: %v", in)

	categoryId, err := database.GetMongoId(in.GetCategoryId())
	if err != nil {
		return &productPb.CategoryResponse{}, err
	}

	// searching in db
	res, err := s.CategoryStore.GetOne(ctx, categoryId)
	if err != nil {
		s.logger.Errorf("error getting category: %v", err)
		return &productPb.CategoryResponse{}, err
	}

	// return response
	s.logger.Infof("get category request finished: %v", res)
	return &productPb.CategoryResponse{
		Message:  "Category found!",
		Category: []*productPb.Category{createCategoryPbDto(res)},
	}, nil
}

func (s *ProductService) UpdateCategory(ctx context.Context, in *productPb.CategoryRequest) (*productPb.CategoryResponse, error) {
	s.logger.Info("update category incoming request: %v", in)

	categoryId, err := database.GetMongoId(in.GetCategoryId())
	if err != nil {
		return &productPb.CategoryResponse{}, err
	}

	_, exists, err := s.CategoryStore.CategoryExists(ctx, categoryId)
	if err != nil {
		s.logger.Errorf("error finding existing category: %v", err)
		return &productPb.CategoryResponse{}, err
	}
	if !exists {
		s.logger.Errorf("error category not exists: %v", errors.New("error category not exists"))
		return &productPb.CategoryResponse{}, errors.New("category not found")
	}
	in.GetCategory().Id = categoryId.Hex()

	// map new info
	categoryToUpdate, err := createCategorySchemaDto(in.GetCategory())
	if err != nil {
		s.logger.Errorf("error creating category schema: %v", err)
		return &productPb.CategoryResponse{}, err
	}

	// update category in DB
	updatedCategory, err := s.CategoryStore.UpdateOne(ctx, categoryToUpdate)
	if err != nil {
		s.logger.Errorf("error updating category in db: %v", err)
		return &productPb.CategoryResponse{}, nil
	}

	// return response
	s.logger.Infof("update category request finished: %v", updatedCategory)
	res := createCategoryPbDto(categoryToUpdate)
	return &productPb.CategoryResponse{
		Message:  "category updated successfully",
		Category: []*productPb.Category{res},
	}, nil
}

func (s *ProductService) DeleteCategory(ctx context.Context, in *productPb.CategoryRequest) (*productPb.CategoryResponse, error) {
	s.logger.Infof("delete category request incoming: %v", in)

	categoryId, err := database.GetMongoId(in.GetCategoryId())
	if err != nil {
		return &productPb.CategoryResponse{}, err
	}

	res, err := s.CategoryStore.DeleteOne(ctx, categoryId)
	if err != nil {
		s.logger.Errorf("error deleting category: %v", err)
		return &productPb.CategoryResponse{}, err
	}

	// return response
	s.logger.Infof("category delete request finished: %v", res)
	return &productPb.CategoryResponse{
		Message:  "category deleted successfully",
		Category: []*productPb.Category{},
	}, nil
}
