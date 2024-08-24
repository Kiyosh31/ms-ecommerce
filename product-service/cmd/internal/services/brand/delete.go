package brand

import (
	"context"

	"github.com/Kiyosh31/ms-ecommerce-common/database"
)

func (s Service) Delete(ctx context.Context, id string) (bool, error) {
	s.logger.Infof("delete brand: %v", id)

	brandId, err := database.GetMongoId(id)
	if err != nil {
		return false, err
	}

	res, err := s.repository.Delete(ctx, brandId)
	if err != nil {
		s.logger.Error(err)
		return false, err
	}

	response := false
	if res.DeletedCount > 0 {
		response = true
	}

	s.logger.Infof("delete brand finished with: %v", id)
	return response, nil
}
