package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/Kiyosh31/ms-ecommerce-common/database"
	inventoryPb "github.com/Kiyosh31/ms-ecommerce/inventory-service/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *InventoryService) CreateInventory(ctx context.Context, in *inventoryPb.InventoryRequest) (*inventoryPb.InventoryResponse, error) {
	s.logger.Infof("Create inventory incoming request: %v", in)

	// create inventory schema for db
	inventoryDto, err := createInventorySchemaDto(in.GetInventory())
	if err != nil {
		s.logger.Errorf("error creating schema: %v", err)
		return &inventoryPb.InventoryResponse{}, err
	}

	// saving inventory to db
	createdInventory, err := s.InventoryStore.CreateOne(ctx, inventoryDto)
	if err != nil {
		s.logger.Errorf("error creating inventory: %v", err)
		return &inventoryPb.InventoryResponse{}, err
	}

	// getting id created in db
	id, ok := createdInventory.InsertedID.(primitive.ObjectID)
	if !ok {
		s.logger.Errorf("error getting id: %v", err)
		return &inventoryPb.InventoryResponse{}, fmt.Errorf("failed to parse _id to string")
	}
	in.GetInventory().Id = id.Hex()

	// return response
	s.logger.Infof("create inventory request finished: %v", createdInventory)
	return &inventoryPb.InventoryResponse{
		Message:   "Inventory created successfully",
		Inventory: []*inventoryPb.Inventory{in.GetInventory()},
	}, nil
}

func (s *InventoryService) GetInventory(ctx context.Context, in *inventoryPb.InventoryRequest) (*inventoryPb.InventoryResponse, error) {
	s.logger.Infof("Get inventory incoming request: %v", in)

	inventoryId, err := database.GetMongoId(in.GetInventoryId())
	if err != nil {
		s.logger.Errorf("error getting inventoryId: %v", err)
		return &inventoryPb.InventoryResponse{}, err
	}

	inventoryFounded, exists, err := s.InventoryStore.InventoryExists(ctx, inventoryId)
	if err != nil {
		s.logger.Errorf("error finding existing inventory: %v", err)
		return &inventoryPb.InventoryResponse{}, err
	}
	if !exists {
		s.logger.Errorf("error inventory not exists: %v", err)
		return &inventoryPb.InventoryResponse{}, errors.New("inventory not exists")
	}

	s.logger.Infof("get inventory request finished: %v", inventoryFounded)
	return &inventoryPb.InventoryResponse{
		Message: "Inventory found successfully",
		Inventory: []*inventoryPb.Inventory{
			{
				Id:        inventoryId.Hex(),
				ProductId: inventoryFounded.ProductID.Hex(),
				Quantity:  inventoryFounded.Quantity,
				Location:  inventoryFounded.Location,
			},
		},
	}, nil
}

func (s *InventoryService) UpdateInventory(ctx context.Context, in *inventoryPb.InventoryRequest) (*inventoryPb.InventoryResponse, error) {
	s.logger.Infof("Update product incoming request: %v", in)

	inventoryId, err := database.GetMongoId(in.GetInventoryId())
	if err != nil {
		s.logger.Errorf("error getting inventoryId: %v", err)
		return &inventoryPb.InventoryResponse{}, err
	}

	_, exists, err := s.InventoryStore.InventoryExists(ctx, inventoryId)
	if err != nil {
		s.logger.Errorf("error finding existing inventory: %v", err)
		return &inventoryPb.InventoryResponse{}, err
	}
	if !exists {
		s.logger.Errorf("error inventory dont exists: %v", err)
		return &inventoryPb.InventoryResponse{}, errors.New("inventory not exist")
	}
	in.GetInventory().Id = inventoryId.Hex()

	inventoryToUpdate, err := createInventorySchemaDto(in.GetInventory())
	if err != nil {
		s.logger.Errorf("error creating inventory schema: %v", err)
		return &inventoryPb.InventoryResponse{}, err
	}

	_, err = s.InventoryStore.UpdateOne(ctx, inventoryToUpdate)
	if err != nil {
		s.logger.Errorf("error updating inventory: %v", err)
		return &inventoryPb.InventoryResponse{}, err
	}

	s.logger.Infof("update inventory request finished: %v", err)
	return &inventoryPb.InventoryResponse{
		Message:   "Inventory updated successfully",
		Inventory: []*inventoryPb.Inventory{in.GetInventory()},
	}, nil
}

func (s *InventoryService) DeleteInventory(ctx context.Context, in *inventoryPb.InventoryRequest) (*inventoryPb.InventoryResponse, error) {
	return &inventoryPb.InventoryResponse{}, nil

}
