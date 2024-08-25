package inventory

import (
	"context"
	"errors"

	"github.com/Kiyosh31/ms-ecommerce-common/database"
	inventoryPb "github.com/Kiyosh31/ms-ecommerce/inventory-service/cmd/proto"
)

func (s *Service) Get(ctx context.Context, id string) (*inventoryPb.Inventory, error) {
	inventoryId, err := database.GetMongoId(id)
	if err != nil {
		return &inventoryPb.Inventory{}, err
	}

	inventory, exist, err := s.repository.Get(ctx, inventoryId)
	if err != nil {
		return &inventoryPb.Inventory{}, err
	}
	if !exist {
		return &inventoryPb.Inventory{}, errors.New("inventory not found")
	}

	res := &inventoryPb.Inventory{
		Id:        inventory.ID.Hex(),
		ProductId: inventory.ProductID.Hex(),
		Quantity:  inventory.Quantity,
		Location:  inventory.Location,
	}

	return res, nil
}
