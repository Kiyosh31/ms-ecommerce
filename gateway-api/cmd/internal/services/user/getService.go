package user

import (
	userPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/generated/user-service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (s Service) GetService() (userPb.UserServiceClient, *grpc.ClientConn) {
	conn, err := grpc.NewClient(s.userClientGrpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		s.logger.Fatalf("Failed to start server: %v", err)
	}

	s.logger.Infof("Dialing user-service at: %v", s.userClientGrpcAddr)
	return userPb.NewUserServiceClient(conn), conn
}
