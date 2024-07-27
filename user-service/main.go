package main

import (
	"log"
	"net"

	"github.com/Kiyosh31/ms-ecommerce/user-service/config"
	"github.com/Kiyosh31/ms-ecommerce/user-service/handler"
	"github.com/Kiyosh31/ms-ecommerce/user-service/service"
	"google.golang.org/grpc"
)

func main() {
	vars, err := config.LoadEnvVars()
	if err != nil {
		log.Fatalf("Could not load env var: %v", err)
	}

	grpServer := grpc.NewServer()

	conn, err := net.Listen("tcp", vars.USER_SERVICE_GRPC_ADDR)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer conn.Close()

	svc := service.NewUserService(vars.USER_SERVICE_GRPC_ADDR)
	handler.NewGrpcUserServiceHandler(grpServer, *svc)

	log.Println("gRPC server started in port: ", vars.USER_SERVICE_GRPC_ADDR)
	if err := grpServer.Serve(conn); err != nil {
		log.Fatal(err.Error())
	}
}
