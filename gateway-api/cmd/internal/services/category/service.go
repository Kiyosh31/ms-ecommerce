package category

import "go.uber.org/zap"

type Service struct {
	categoryClientGrpcAddr string
	logger                 *zap.SugaredLogger
}

func NewCategoryService(
	categoryClientGrpcAddr string,
	logger *zap.SugaredLogger,
) *Service {
	return &Service{
		categoryClientGrpcAddr: categoryClientGrpcAddr,
		logger:                 logger,
	}
}
