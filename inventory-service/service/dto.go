package service

import (
	"github.com/Kiyosh31/ms-ecommerce-common/database"
	"github.com/Kiyosh31/ms-ecommerce/inventory-service/inventory_types"
	inventoryPb "github.com/Kiyosh31/ms-ecommerce/inventory-service/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func createInventorySchemaDto(in *inventoryPb.Inventory) (inventory_types.InventorySchema, error) {
	var inventoryId primitive.ObjectID
	var err error

	if in.GetId() != "" {
		inventoryId, err = database.GetMongoId(in.GetId())
		if err != nil {
			return inventory_types.InventorySchema{}, err
		}
	}

	productId, err := database.GetMongoId(in.GetProductId())
	if err != nil {
		return inventory_types.InventorySchema{}, err
	}

	return inventory_types.InventorySchema{
		ID:        inventoryId,
		ProductID: productId,
		Quantity:  in.GetQuantity(),
		Location:  in.GetLocation(),
	}, nil
}

func createInventoryResponseDto() {

}
