package category

import (
	"github.com/Kiyosh31/ms-ecommerce/product-service/cmd/internal/ports"
	"github.com/Kiyosh31/ms-ecommerce/product-service/cmd/internal/services/category"
	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/cmd/proto"
	"go.uber.org/zap"
)

type Handler struct {
	productPb.UnimplementedCategoryServiceServer
	categoryService ports.CategoryService
	logger          *zap.SugaredLogger
}

func NewCategoryHandler(
	categoryService *category.Service,
	logger *zap.SugaredLogger,
) *Handler {
	return &Handler{
		categoryService: categoryService,
		logger:          logger,
	}
}
