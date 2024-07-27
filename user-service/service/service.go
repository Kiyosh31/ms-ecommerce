package service

type UserService struct {
	grpcAddr string
}

func NewUserService(grpcAddr string) *UserService {
	return &UserService{grpcAddr: grpcAddr}
}
