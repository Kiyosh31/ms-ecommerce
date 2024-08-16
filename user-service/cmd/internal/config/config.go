package config

import "github.com/Kiyosh31/ms-ecommerce-common/utils"

type envVars struct {
	USER_SERVICE_GRPC_ADDR           string
	DB_CONNECTION_LINK               string
	USER_SERVICE_DATABASE_NAME       string
	USER_SERVICE_DATABASE_COLLECTION string
}

func LoadEnvVars() (envVars, error) {
	USER_SERVICE_GRPC_ADDR, err := utils.GetEnvVar("USER_SERVICE_GRPC_ADDR")
	if err != nil {
		return envVars{}, err
	}

	USER_SERVICE_DATABASE_NAME, err := utils.GetEnvVar("USER_SERVICE_DATABASE_NAME")
	if err != nil {
		return envVars{}, err
	}

	DB_CONNECTION_LINK, err := utils.GetEnvVar("DB_CONNECTION_LINK")
	if err != nil {
		return envVars{}, err
	}

	USER_SERVICE_DATABASE_COLLECTION, err := utils.GetEnvVar("USER_SERVICE_DATABASE_COLLECTION")
	if err != nil {
		return envVars{}, err
	}

	return envVars{
		USER_SERVICE_GRPC_ADDR:           USER_SERVICE_GRPC_ADDR,
		USER_SERVICE_DATABASE_NAME:       USER_SERVICE_DATABASE_NAME,
		DB_CONNECTION_LINK:               DB_CONNECTION_LINK,
		USER_SERVICE_DATABASE_COLLECTION: USER_SERVICE_DATABASE_COLLECTION,
	}, nil
}
