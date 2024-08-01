package service

import (
	userPb "github.com/Kiyosh31/ms-ecommerce/user-service/proto"
	"github.com/Kiyosh31/ms-ecommerce/user-service/store"
	"go.uber.org/zap"
)

type UserService struct {
	userPb.UnimplementedUserServiceServer
	GrpcAddr  string
	UserStore store.UserStore
	logger    *zap.SugaredLogger
}

func NewUserService(grpcAddr string, userStore store.UserStore, logger *zap.SugaredLogger) *UserService {
	return &UserService{
		GrpcAddr:  grpcAddr,
		UserStore: userStore,
		logger:    logger,
	}
}
