package service

import (
	"github.com/Kiyosh31/ms-ecommerce-common/database"

	"github.com/Kiyosh31/ms-ecommerce/product-service/product_types"
	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func createProductSchemaDto(in *productPb.Product) (product_types.ProductSchema, error) {
	var productId primitive.ObjectID
	var err error

	if in.GetId() != "" {
		productId, err = database.GetMongoId(in.GetId())
		if err != nil {
			return product_types.ProductSchema{}, err
		}
	}

	category, err := createCategorySchemaDto(in.GetCategory())
	if err != nil {
		return product_types.ProductSchema{}, err
	}

	brand, err := createBrandSchema(in.GetBrand())
	if err != nil {
		return product_types.ProductSchema{}, err
	}

	return product_types.ProductSchema{
		ID:          productId,
		Name:        in.GetName(),
		Description: in.GetDescription(),
		Price:       float32(in.GetPrice()),
		Category:    category,
		Brand:       brand,
		Images:      in.GetImages(),
		Attributes:  createAttributesDto(in.GetAttributes()),
		Inventory:   createInventoryDto(in.GetInventory()),
	}, nil
}

func createCategorySchemaDto(in *productPb.Category) (product_types.CategorySchema, error) {
	var categoryId primitive.ObjectID
	var err error

	if in.GetId() != "" {
		categoryId, err = database.GetMongoId(in.GetId())
		if err != nil {
			return product_types.CategorySchema{}, err
		}
	}

	return product_types.CategorySchema{
		ID:          categoryId,
		Name:        in.GetName(),
		Description: in.GetDescription(),
	}, nil
}

func createBrandSchema(in *productPb.Brand) (product_types.BrandSchema, error) {
	var brandId primitive.ObjectID
	var err error

	if in.GetId() != "" {
		brandId, err = database.GetMongoId(in.GetId())
		if err != nil {
			return product_types.BrandSchema{}, err
		}
	}

	return product_types.BrandSchema{
		ID:          brandId,
		Name:        in.GetName(),
		Description: in.GetDescription(),
	}, nil
}

func createAttributesDto(in *productPb.Attributes) product_types.Attributes {
	return product_types.Attributes{
		Size:     in.GetSize(),
		Color:    in.GetColor(),
		Material: in.GetMaterial(),
	}
}

func createInventoryDto(in *productPb.Inventory) product_types.Inventory {
	return product_types.Inventory{}
}

func createProductResponseDto(message string, in product_types.ProductSchema) *productPb.ProductResponse {
	return &productPb.ProductResponse{
		Message: message,
		Product: []*productPb.Product{
			&productPb.Product{
				Id:          in.ID.Hex(),
				Name:        in.Name,
				Description: in.Description,
				Price:       in.Price,
				Category:    createCategoryPbDto(in.Category),
				Brand:       createBrandPbDto(in.Brand),
				Images:      in.Images,
				Attributes:  createAttributesPbDto(in.Attributes),
				Inventory:   createInventoryPbDto(in.Inventory),
			},
		},
	}
}

func createCategoryPbDto(in product_types.CategorySchema) *productPb.Category {
	return &productPb.Category{
		Id:          in.ID.Hex(),
		Name:        in.Name,
		Description: in.Description,
	}
}

func createBrandPbDto(in product_types.BrandSchema) *productPb.Brand {
	return &productPb.Brand{
		Id:          in.ID.Hex(),
		Name:        in.Name,
		Description: in.Description,
	}
}

func createAttributesPbDto(in product_types.Attributes) *productPb.Attributes {
	return &productPb.Attributes{
		Size:     in.Size,
		Color:    in.Color,
		Material: in.Material,
	}
}

func createInventoryPbDto(in product_types.Inventory) *productPb.Inventory {
	return &productPb.Inventory{}
}

func createMultipleProductsResponseDto(message string, in []product_types.ProductSchema) *productPb.ProductResponse {
	var products []*productPb.Product

	for _, product := range in {
		prod := productPb.Product{
			Id:          product.ID.Hex(),
			Name:        product.Name,
			Price:       product.Price,
			Description: product.Description,
			Category:    createCategoryPbDto(product.Category),
			Brand:       createBrandPbDto(product.Brand),
			Images:      product.Images,
			Attributes:  createAttributesPbDto(product.Attributes),
			Inventory:   createInventoryPbDto(product.Inventory),
		}

		products = append(products, &prod)
	}

	return &productPb.ProductResponse{
		Message: message,
		Product: products,
	}
}
