package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/Kiyosh31/ms-ecommerce-common/database"
	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *ProductService) CreateBrand(ctx context.Context, in *productPb.BrandRequest) (*productPb.BrandResponse, error) {
	s.logger.Info("create brand incoming request: %v", in)

	_, exists, err := s.BrandStore.GetOneByName(ctx, in.GetBrand().GetName())
	if err != nil {
		s.logger.Errorf("error finding existing brand: %v", err)
		return &productPb.BrandResponse{}, err
	}
	if exists {
		s.logger.Errorf("error brand already exists: %v", err)
		return &productPb.BrandResponse{}, errors.New("brand already exists")
	}

	// create schema for DB
	brandDto, err := createBrandSchemaDto(in.GetBrand())
	if err != nil {
		s.logger.Errorf("error creating category schema: %v", err)
		return &productPb.BrandResponse{}, nil
	}

	// saving category to db
	createdBrand, err := s.BrandStore.CreateOne(ctx, brandDto)
	if err != nil {
		s.logger.Errorf("error creating category in db: %v", err)
		return &productPb.BrandResponse{}, nil
	}

	// getting categoryID created in db
	id, ok := createdBrand.InsertedID.(primitive.ObjectID)
	if !ok {
		s.logger.Errorf("error getting id: %v", err)
		return &productPb.BrandResponse{}, fmt.Errorf("failed to parse _id to string")
	}
	in.GetBrand().Id = id.Hex()

	// return response
	s.logger.Info("create brand request finished: %v", createdBrand)
	return &productPb.BrandResponse{
		Message: "Category created successfully",
		Brand:   []*productPb.Brand{in.GetBrand()},
	}, nil
}

func (s *ProductService) GetBrand(ctx context.Context, in *productPb.BrandRequest) (*productPb.BrandResponse, error) {
	s.logger.Info("get brand incoming request: %v", in)

	brandId, err := database.GetMongoId(in.GetBrandId())
	if err != nil {
		return &productPb.BrandResponse{}, err
	}

	// searching in db
	res, err := s.BrandStore.GetOne(ctx, brandId)
	if err != nil {
		s.logger.Errorf("error getting brand: %v", err)
		return &productPb.BrandResponse{}, err
	}

	// return response
	s.logger.Infof("get brand request finished: %v", res)
	return &productPb.BrandResponse{
		Message: "Category found!",
		Brand:   []*productPb.Brand{createBrandPbDto(res)},
	}, nil

}

func (s *ProductService) UpdateBrand(ctx context.Context, in *productPb.BrandRequest) (*productPb.BrandResponse, error) {
	s.logger.Info("update brand incoming request: %v", in)

	brandId, err := database.GetMongoId(in.GetBrandId())
	if err != nil {
		return &productPb.BrandResponse{}, err
	}

	_, exists, err := s.BrandStore.BrandExists(ctx, brandId)
	if err != nil {
		s.logger.Errorf("error finding existing brand: %v", err)
		return &productPb.BrandResponse{}, err
	}
	if !exists {
		s.logger.Errorf("error brand not exists: %v", errors.New("error brand not exists"))
		return &productPb.BrandResponse{}, errors.New("brand not found")
	}
	in.GetBrand().Id = brandId.Hex()

	// map new info
	brandToUpdate, err := createBrandSchemaDto(in.GetBrand())
	if err != nil {
		s.logger.Errorf("error creating brand schema: %v", err)
		return &productPb.BrandResponse{}, err
	}

	// update in DB
	updatedBrand, err := s.BrandStore.UpdateOne(ctx, brandToUpdate)
	if err != nil {
		s.logger.Errorf("error updating brand in db: %v", err)
		return &productPb.BrandResponse{}, nil
	}

	// return response
	s.logger.Infof("update brand request finished: %v", updatedBrand)
	res := createBrandPbDto(brandToUpdate)
	return &productPb.BrandResponse{
		Message: "category updated successfully",
		Brand:   []*productPb.Brand{res},
	}, nil
}

func (s *ProductService) DeleteBrand(ctx context.Context, in *productPb.BrandRequest) (*productPb.BrandResponse, error) {
	s.logger.Infof("delete brand request incoming: %v", in)

	brandId, err := database.GetMongoId(in.GetBrandId())
	if err != nil {
		return &productPb.BrandResponse{}, err
	}

	res, err := s.BrandStore.DeleteOne(ctx, brandId)
	if err != nil {
		s.logger.Errorf("error deleting brand: %v", err)
		return &productPb.BrandResponse{}, err
	}

	// return response
	s.logger.Infof("brand delete request finished: %v", res)
	return &productPb.BrandResponse{
		Message: "category deleted successfully",
		Brand:   []*productPb.Brand{},
	}, nil
}
