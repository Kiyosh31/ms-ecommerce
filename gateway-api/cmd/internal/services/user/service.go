package user

import "go.uber.org/zap"

type Service struct {
	userClientGrpcAddr string
	logger             *zap.SugaredLogger
}

func NewUserService(userClientGrpcAddr string, logger *zap.SugaredLogger) *Service {
	return &Service{
		userClientGrpcAddr: userClientGrpcAddr,
		logger:             logger,
	}
}
