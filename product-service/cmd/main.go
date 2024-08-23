package main

import (
	"log"
	"net"

	customlogger "github.com/Kiyosh31/ms-ecommerce-common/custom_logger"
	"github.com/Kiyosh31/ms-ecommerce-common/database"
	"github.com/Kiyosh31/ms-ecommerce/product-service/cmd/api/handlers/brand"
	"github.com/Kiyosh31/ms-ecommerce/product-service/cmd/api/handlers/category"
	productHandler "github.com/Kiyosh31/ms-ecommerce/product-service/cmd/api/handlers/product"
	"github.com/Kiyosh31/ms-ecommerce/product-service/cmd/internal/config"
	brandRepo "github.com/Kiyosh31/ms-ecommerce/product-service/cmd/internal/repositories/mongo/brand"
	categoryRepo "github.com/Kiyosh31/ms-ecommerce/product-service/cmd/internal/repositories/mongo/category"
	productRepo "github.com/Kiyosh31/ms-ecommerce/product-service/cmd/internal/repositories/mongo/product"
	brandService "github.com/Kiyosh31/ms-ecommerce/product-service/cmd/internal/services/brand"
	categoryService "github.com/Kiyosh31/ms-ecommerce/product-service/cmd/internal/services/category"
	productService "github.com/Kiyosh31/ms-ecommerce/product-service/cmd/internal/services/product"
	productPb "github.com/Kiyosh31/ms-ecommerce/product-service/cmd/proto"
	"google.golang.org/grpc"
)

func main() {
	logger, err := customlogger.InitLogger()
	if err != nil {
		log.Fatalf("error logger init: %v", err)
	}

	vars, err := config.LoadEnvVars()
	if err != nil {
		logger.Fatalf("Could not load env var: %v", err)
	}

	grpServer := grpc.NewServer()

	conn, err := net.Listen("tcp", vars.PRODUCT_SERVICE_GRPC_ADDR)
	if err != nil {
		logger.Fatalf("Failed to listen: %v", err)
	}
	defer conn.Close()

	mongoClient, err := database.ConnectToDB(vars.DB_CONNECTION_LINK)
	if err != nil {
		logger.Fatalf("Could not connect to database: %v", err)
	}
	defer database.DisconnectOfDB(mongoClient)

	categoryRepository := categoryRepo.NewCategoryRepository(
		mongoClient,
		vars.PRODUCT_SERVICE_DATABASE_NAME,
		vars.PRODUCT_SERVICE_CATEGORIES_DATABASE_COLLECTION,
	)
	categoryService := categoryService.NewCategoryService(
		categoryRepository,
		logger,
	)
	categoryHandler := category.NewCategoryHandler(
		categoryService,
		logger,
	)

	brandRepository := brandRepo.NewBrandRepository(
		mongoClient,
		vars.PRODUCT_SERVICE_DATABASE_NAME,
		vars.PRODUCT_SERVICE_BRANDS_DATABASE_COLLECTION,
	)
	brandService := brandService.NewBrandService(
		brandRepository,
		logger,
	)
	brandHandler := brand.NewBrandHandler(
		brandService,
		logger,
	)

	productRepository := productRepo.NewProductRepository(
		mongoClient,
		vars.PRODUCT_SERVICE_DATABASE_NAME,
		vars.PRODUCT_SERVICE_PRODUCTS_DATABASE_COLLECTION,
	)
	productService := productService.NewProductService(
		productRepository,
		logger,
	)
	productHandler := productHandler.NewProductHandler(
		productService,
		logger,
	)

	productPb.RegisterProductServiceServer(grpServer, productHandler)
	productPb.RegisterCategoryServiceServer(grpServer, categoryHandler)
	productPb.RegisterBrandServiceServer(grpServer, brandHandler)

	// productStore := store.NewProductStore(
	// 	mongoClient,
	// 	vars.PRODUCT_SERVICE_DATABASE_NAME,
	// 	vars.PRODUCT_SERVICE_PRODUCTS_DATABASE_COLLECTION,
	// )
	// brandStore := store.NewBrandStore(
	// 	mongoClient,
	// 	vars.PRODUCT_SERVICE_DATABASE_NAME,
	// 	vars.PRODUCT_SERVICE_BRANDS_DATABASE_COLLECTION,
	// )
	// categoryStore := store.NewCategoryStore(
	// 	mongoClient,
	// 	vars.PRODUCT_SERVICE_DATABASE_NAME,
	// 	vars.PRODUCT_SERVICE_CATEGORIES_DATABASE_COLLECTION,
	// )
	// svc := service.NewProductService(
	// 	vars.PRODUCT_SERVICE_GRPC_ADDR,
	// 	*productStore,
	// 	*brandStore,
	// 	*categoryStore,
	// 	logger,
	// )
	// productPb.RegisterProductServiceServer(grpServer, svc)

	logger.Infof("gRPC server started in port: %v", vars.PRODUCT_SERVICE_GRPC_ADDR)
	if err := grpServer.Serve(conn); err != nil {
		logger.Fatal(err.Error())
	}
}
