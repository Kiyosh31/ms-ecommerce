package service

type UserService struct {
	addr string
}

func NewUserService(addr string) *UserService {
	return &UserService{addr: addr}
}

func (s *UserService) Run(grpcAddr string) {

}
