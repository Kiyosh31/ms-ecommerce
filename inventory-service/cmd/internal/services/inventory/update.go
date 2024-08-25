package inventory

import (
	"context"
	"errors"

	"github.com/Kiyosh31/ms-ecommerce/inventory-service/cmd/internal/domain"
	inventoryPb "github.com/Kiyosh31/ms-ecommerce/inventory-service/cmd/proto"
)

func (s Service) Update(ctx context.Context, inventory domain.InventorySchema) (*inventoryPb.Inventory, error) {
	// validate user exists
	_, exist, err := s.repository.Get(ctx, inventory.ID)
	if err != nil {
		return &inventoryPb.Inventory{}, err
	}
	if !exist {
		return &inventoryPb.Inventory{}, errors.New("inventory not found")
	}

	//update to db
	_, err = s.repository.Update(ctx, inventory)
	if err != nil {
		return &inventoryPb.Inventory{}, err
	}

	// translate response
	res := &inventoryPb.Inventory{
		Id:        inventory.ID.Hex(),
		ProductId: inventory.ProductID.Hex(),
		Quantity:  inventory.Quantity,
		Location:  inventory.Location,
	}

	return res, nil
}
