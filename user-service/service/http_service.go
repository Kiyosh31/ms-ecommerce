package service

import "net/http"

type UserServiceHttp struct {
	addr string
}

func NewUserService(addr string) *UserServiceHttp {
	return &UserServiceHttp{addr: addr}
}

func (s *UserServiceHttp) Run() error {
	router := http.NewServeMux()

	return nil
}
