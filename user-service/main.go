package main

import (
	"log"
	"net"

	"github.com/Kiyosh31/ms-ecommerce-common/database"
	"github.com/Kiyosh31/ms-ecommerce/user-service/config"
	"github.com/Kiyosh31/ms-ecommerce/user-service/handler"
	"github.com/Kiyosh31/ms-ecommerce/user-service/service"
	"github.com/Kiyosh31/ms-ecommerce/user-service/store"
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

	mongoClient, err := database.ConnectToDB(vars.DB_CONNECTION_LINK)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer database.DisconnectOfDB(mongoClient)

	userStore := store.NewUserStore(mongoClient, vars.USER_SERVICE_DATABASE_NAME, vars.USER_SERVICE_DATABASE_COLLECTION)
	svc := service.NewUserService(vars.USER_SERVICE_GRPC_ADDR, *userStore)
	handler.NewGrpcUserServiceHandler(grpServer, *svc)

	log.Println("gRPC server started in port: ", vars.USER_SERVICE_GRPC_ADDR)
	if err := grpServer.Serve(conn); err != nil {
		log.Fatal(err.Error())
	}
}
