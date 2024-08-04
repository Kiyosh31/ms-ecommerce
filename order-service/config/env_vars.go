package config

import "github.com/Kiyosh31/ms-ecommerce-common/utils"

type envVars struct {
	ORDER_SERVICE_GRPC_ADDR           string
	DB_CONNECTION_LINK                string
	ORDER_SERVICE_DATABASE_NAME       string
	ORDER_SERVICE_DATABASE_COLLECTION string
}

func LoadEnvVars() (envVars, error) {

	ORDER_SERVICE_GRPC_ADDR, err := utils.GetEnvVar("ORDER_SERVICE_GRPC_ADDR")
	if err != nil {
		return envVars{}, err
	}

	ORDER_SERVICE_DATABASE_NAME, err := utils.GetEnvVar("ORDER_SERVICE_DATABASE_NAME")
	if err != nil {
		return envVars{}, err
	}

	DB_CONNECTION_LINK, err := utils.GetEnvVar("DB_CONNECTION_LINK")
	if err != nil {
		return envVars{}, err
	}

	ORDER_SERVICE_DATABASE_COLLECTION, err := utils.GetEnvVar("ORDER_SERVICE_DATABASE_COLLECTION")
	if err != nil {
		return envVars{}, err
	}

	return envVars{
		ORDER_SERVICE_GRPC_ADDR:           ORDER_SERVICE_GRPC_ADDR,
		ORDER_SERVICE_DATABASE_NAME:       ORDER_SERVICE_DATABASE_NAME,
		DB_CONNECTION_LINK:                DB_CONNECTION_LINK,
		ORDER_SERVICE_DATABASE_COLLECTION: ORDER_SERVICE_DATABASE_COLLECTION,
	}, nil
}
