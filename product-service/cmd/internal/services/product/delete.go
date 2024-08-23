package product

import "context"

func (s Service) Delete(ctx context.Context, id string) (bool, error) {
	return true, nil
}
