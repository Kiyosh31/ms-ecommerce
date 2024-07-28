package service

import "github.com/Kiyosh31/ms-ecommerce/user-service/store"

type UserService struct {
	GrpcAddr  string
	UserStore store.UserStore
}

func NewUserService(grpcAddr string, userStore store.UserStore) *UserService {
	return &UserService{
		GrpcAddr:  grpcAddr,
		UserStore: userStore,
	}
}
