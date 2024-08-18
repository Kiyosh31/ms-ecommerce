package product

import "go.uber.org/zap"

type Service struct {
	productClientGrpcAddr string
	logger                *zap.SugaredLogger
}

func NewProductService(productClientGrpcAddr string,
	logger *zap.SugaredLogger) *Service {
	return &Service{
		productClientGrpcAddr: productClientGrpcAddr,
		logger:                logger,
	}
}
