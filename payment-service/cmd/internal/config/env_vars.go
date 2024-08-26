package config

import "github.com/Kiyosh31/ms-ecommerce-common/utils"

type envVars struct {
	PAYMENT_SERVICE_GRPC_ADDR           string
	DB_CONNECTION_LINK                  string
	PAYMENT_SERVICE_DATABASE_NAME       string
	PAYMENT_SERVICE_DATABASE_COLLECTION string
}

func LoadEnvVars() (*envVars, error) {
	PAYMENT_SERVICE_GRPC_ADDR, err := utils.GetEnvVar("PAYMENT_SERVICE_GRPC_ADDR")
	if err != nil {
		return &envVars{}, err
	}

	PAYMENT_SERVICE_DATABASE_NAME, err := utils.GetEnvVar("PAYMENT_SERVICE_DATABASE_NAME")
	if err != nil {
		return &envVars{}, err
	}

	DB_CONNECTION_LINK, err := utils.GetEnvVar("DB_CONNECTION_LINK")
	if err != nil {
		return &envVars{}, err
	}

	PAYMENT_SERVICE_DATABASE_COLLECTION, err := utils.GetEnvVar("PAYMENT_SERVICE_DATABASE_COLLECTION")
	if err != nil {
		return &envVars{}, err
	}

	return &envVars{
		PAYMENT_SERVICE_GRPC_ADDR:           PAYMENT_SERVICE_GRPC_ADDR,
		PAYMENT_SERVICE_DATABASE_NAME:       PAYMENT_SERVICE_DATABASE_NAME,
		DB_CONNECTION_LINK:                  DB_CONNECTION_LINK,
		PAYMENT_SERVICE_DATABASE_COLLECTION: PAYMENT_SERVICE_DATABASE_COLLECTION,
	}, nil
}
