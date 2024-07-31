package main

import (
	"log"
	"net"

	"github.com/Kiyosh31/ms-ecommerce-common/database"
	"github.com/Kiyosh31/ms-ecommerce/cart-service/config"
	cartPb "github.com/Kiyosh31/ms-ecommerce/cart-service/proto"
	"github.com/Kiyosh31/ms-ecommerce/cart-service/service"
	"github.com/Kiyosh31/ms-ecommerce/cart-service/store"
	"google.golang.org/grpc"
)

func main() {
	vars, err := config.LoadEnvVars()
	if err != nil {
		log.Fatalf("Could not load env var: %v", err)
	}

	grpServer := grpc.NewServer()

	conn, err := net.Listen("tcp", vars.CART_SERVICE_GRPC_ADDR)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer conn.Close()

	mongoClient, err := database.ConnectToDB(vars.DB_CONNECTION_LINK)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer database.DisconnectOfDB(mongoClient)

	cartStore := store.NewCartStore(mongoClient, vars.CART_SERVICE_DATABASE_NAME, vars.CART_SERVICE_DATABASE_COLLECTION)
	svc := service.NewCartService(vars.CART_SERVICE_GRPC_ADDR, *cartStore)
	cartPb.RegisterCartServiceServer(grpServer, svc)

	log.Println("gRPC server started in port: ", vars.CART_SERVICE_GRPC_ADDR)
	if err := grpServer.Serve(conn); err != nil {
		log.Fatal(err.Error())
	}
}
