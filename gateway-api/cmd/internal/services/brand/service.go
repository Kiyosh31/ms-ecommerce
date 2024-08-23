package brand

import "go.uber.org/zap"

type Service struct {
	brandClientGrpcAddr string
	logger              *zap.SugaredLogger
}

func NewBrandService(
	brandClientGrpcAddr string,
	logger *zap.SugaredLogger,
) *Service {
	return &Service{
		brandClientGrpcAddr: brandClientGrpcAddr,
		logger:              logger,
	}
}
