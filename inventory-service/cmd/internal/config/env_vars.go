package config

import "github.com/Kiyosh31/ms-ecommerce-common/utils"

type envVars struct {
	INVENTORY_SERVICE_GRPC_ADDR           string
	DB_CONNECTION_LINK                    string
	INVENTORY_SERVICE_DATABASE_NAME       string
	INVENTORY_SERVICE_DATABASE_COLLECTION string
}

func LoadEnvVars() (*envVars, error) {
	INVENTORY_SERVICE_GRPC_ADDR, err := utils.GetEnvVar("INVENTORY_SERVICE_GRPC_ADDR")
	if err != nil {
		return &envVars{}, err
	}

	INVENTORY_SERVICE_DATABASE_NAME, err := utils.GetEnvVar("INVENTORY_SERVICE_DATABASE_NAME")
	if err != nil {
		return &envVars{}, err
	}

	DB_CONNECTION_LINK, err := utils.GetEnvVar("DB_CONNECTION_LINK")
	if err != nil {
		return &envVars{}, err
	}

	INVENTORY_SERVICE_DATABASE_COLLECTION, err := utils.GetEnvVar("INVENTORY_SERVICE_DATABASE_COLLECTION")
	if err != nil {
		return &envVars{}, err
	}

	return &envVars{
		INVENTORY_SERVICE_GRPC_ADDR:           INVENTORY_SERVICE_GRPC_ADDR,
		DB_CONNECTION_LINK:                    DB_CONNECTION_LINK,
		INVENTORY_SERVICE_DATABASE_NAME:       INVENTORY_SERVICE_DATABASE_NAME,
		INVENTORY_SERVICE_DATABASE_COLLECTION: INVENTORY_SERVICE_DATABASE_COLLECTION,
	}, nil
}
