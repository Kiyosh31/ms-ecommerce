package ports

import (
	"context"

	"github.com/Kiyosh31/ms-ecommerce/inventory-service/cmd/internal/domain"
	inventoryPb "github.com/Kiyosh31/ms-ecommerce/inventory-service/cmd/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type InventoryService interface {
	Get(ctx context.Context, id string) (*inventoryPb.Inventory, error)
	Update(ctx context.Context, inventoryToUpdate domain.InventorySchema) (*inventoryPb.Inventory, error)
}

type InventoryRepository interface {
	Get(ctx context.Context, id primitive.ObjectID) (domain.InventorySchema, bool, error)
	Update(ctx context.Context, inventory domain.InventorySchema) (mongo.UpdateResult, error)
}
