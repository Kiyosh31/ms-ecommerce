package config

import "github.com/Kiyosh31/ms-ecommerce-common/utils"

type EnvVars struct {
	PRODUCT_SERVICE_GRPC_ADDR           string
	DB_CONNECTION_LINK                  string
	PRODUCT_SERVICE_DATABASE_NAME       string
	PRODUCT_SERVICE_DATABASE_COLLECTION string
}

func LoadEnvVars() (EnvVars, error) {
	PRODUCT_SERVICE_GRPC_ADDR, err := utils.GetEnvVar("PRODUCT_SERVICE_GRPC_ADDR")
	if err != nil {
		return EnvVars{}, err
	}

	DB_CONNECTION_LINK, err := utils.GetEnvVar("DB_CONNECTION_LINK")
	if err != nil {
		return EnvVars{}, err
	}

	PRODUCT_SERVICE_DATABASE_NAME, err := utils.GetEnvVar("PRODUCT_SERVICE_DATABASE_NAME")
	if err != nil {
		return EnvVars{}, err
	}

	PRODUCT_SERVICE_DATABASE_COLLECTION, err := utils.GetEnvVar("PRODUCT_SERVICE_DATABASE_COLLECTION")
	if err != nil {
		return EnvVars{}, err
	}

	return EnvVars{
		PRODUCT_SERVICE_GRPC_ADDR:           PRODUCT_SERVICE_GRPC_ADDR,
		DB_CONNECTION_LINK:                  DB_CONNECTION_LINK,
		PRODUCT_SERVICE_DATABASE_NAME:       PRODUCT_SERVICE_DATABASE_NAME,
		PRODUCT_SERVICE_DATABASE_COLLECTION: PRODUCT_SERVICE_DATABASE_COLLECTION,
	}, nil
}
