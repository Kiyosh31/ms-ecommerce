package config

import "github.com/Kiyosh31/ms-ecommerce-common/utils"

type envVars struct {
	CART_SERVICE_GRPC_ADDR           string
	DB_CONNECTION_LINK               string
	CART_SERVICE_DATABASE_NAME       string
	CART_SERVICE_DATABASE_COLLECTION string
}

func LoadEnvVars() (envVars, error) {

	CART_SERVICE_GRPC_ADDR, err := utils.GetEnvVar("CART_SERVICE_GRPC_ADDR")
	if err != nil {
		return envVars{}, err
	}

	CART_SERVICE_DATABASE_NAME, err := utils.GetEnvVar("CART_SERVICE_DATABASE_NAME")
	if err != nil {
		return envVars{}, err
	}

	DB_CONNECTION_LINK, err := utils.GetEnvVar("DB_CONNECTION_LINK")
	if err != nil {
		return envVars{}, err
	}

	CART_SERVICE_DATABASE_COLLECTION, err := utils.GetEnvVar("CART_SERVICE_DATABASE_COLLECTION")
	if err != nil {
		return envVars{}, err
	}

	return envVars{
		CART_SERVICE_GRPC_ADDR:           CART_SERVICE_GRPC_ADDR,
		CART_SERVICE_DATABASE_NAME:       CART_SERVICE_DATABASE_NAME,
		DB_CONNECTION_LINK:               DB_CONNECTION_LINK,
		CART_SERVICE_DATABASE_COLLECTION: CART_SERVICE_DATABASE_COLLECTION,
	}, nil
}
