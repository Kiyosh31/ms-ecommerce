package main

import (
	"log"
	"net"

	customlogger "github.com/Kiyosh31/ms-ecommerce-common/custom_logger"
	"github.com/Kiyosh31/ms-ecommerce-common/database"

	userHandler "github.com/Kiyosh31/ms-ecommerce/user-service/cmd/api/handlers/user"
	"github.com/Kiyosh31/ms-ecommerce/user-service/cmd/internal/config"
	userRepo "github.com/Kiyosh31/ms-ecommerce/user-service/cmd/internal/repositories/mongo/user"
	userService "github.com/Kiyosh31/ms-ecommerce/user-service/cmd/internal/services/user"
	userPb "github.com/Kiyosh31/ms-ecommerce/user-service/cmd/proto"

	"google.golang.org/grpc"
)

func main() {
	logger, err := customlogger.InitLogger()
	if err != nil {
		log.Fatalf("error init logger: %v", err)
	}
	defer logger.Sync()

	vars, err := config.LoadEnvVars()
	if err != nil {
		logger.Fatalf("Could not load env var: %v", err)
	}

	grpcServer := grpc.NewServer()

	conn, err := net.Listen("tcp", vars.USER_SERVICE_GRPC_ADDR)
	if err != nil {
		logger.Fatalf("Failed to listen: %v", err)
	}
	defer conn.Close()

	client, err := database.ConnectToDB(vars.DB_CONNECTION_LINK)
	if err != nil {
		logger.Fatalf("Could not connect to database: %v", err)
	}
	defer database.DisconnectOfDB(client)

	userRepository := userRepo.NewUserRepository(client, vars.USER_SERVICE_DATABASE_NAME, vars.USER_SERVICE_DATABASE_COLLECTION)
	userService := userService.NewUserService(userRepository, logger, vars.SECRET_KEY, vars.TOKEN_DURATION_TIME_MINUTES)
	userHandler := userHandler.NewUserHandler(userService, logger)
	userPb.RegisterUserServiceServer(grpcServer, userHandler)

	logger.Infof("gRPC server started in port: %v", vars.USER_SERVICE_GRPC_ADDR)
	if err := grpcServer.Serve(conn); err != nil {
		logger.Fatal(err.Error())
	}
}
