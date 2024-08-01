package config

import "github.com/Kiyosh31/ms-ecommerce-common/utils"

type envVars struct {
	STOCK_SERVICE_HTTP_ADDR           string
	DB_CONNECTION_LINK                string
	STOCK_SERVICE_DATABASE_NAME       string
	STOCK_SERVICE_DATABASE_COLLECTION string
}

func LoadEnvVars() (envVars, error) {

	STOCK_SERVICE_HTTP_ADDR, err := utils.GetEnvVar("STOCK_SERVICE_HTTP_ADDR")
	if err != nil {
		return envVars{}, err
	}

	STOCK_SERVICE_DATABASE_NAME, err := utils.GetEnvVar("STOCK_SERVICE_DATABASE_NAME")
	if err != nil {
		return envVars{}, err
	}

	STOCK_SERVICE_DATABASE_COLLECTION, err := utils.GetEnvVar("STOCK_SERVICE_DATABASE_COLLECTION")
	if err != nil {
		return envVars{}, err
	}

	DB_CONNECTION_LINK, err := utils.GetEnvVar("DB_CONNECTION_LINK")
	if err != nil {
		return envVars{}, err
	}

	return envVars{
		STOCK_SERVICE_HTTP_ADDR:           STOCK_SERVICE_HTTP_ADDR,
		DB_CONNECTION_LINK:                DB_CONNECTION_LINK,
		STOCK_SERVICE_DATABASE_NAME:       STOCK_SERVICE_DATABASE_NAME,
		STOCK_SERVICE_DATABASE_COLLECTION: STOCK_SERVICE_DATABASE_COLLECTION,
	}, nil
}
