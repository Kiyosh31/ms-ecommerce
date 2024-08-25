package config

import (
	"fmt"

	"github.com/Kiyosh31/ms-ecommerce-common/utils"
)

type envVars struct {
	USER_SERVICE_GRPC_ADDR           string
	DB_CONNECTION_LINK               string
	USER_SERVICE_DATABASE_NAME       string
	USER_SERVICE_DATABASE_COLLECTION string
	SECRET_KEY                       string
	TOKEN_DURATION_TIME              string
}

const minSecretKeyLength = 32

func LoadEnvVars() (*envVars, error) {
	USER_SERVICE_GRPC_ADDR, err := utils.GetEnvVar("USER_SERVICE_GRPC_ADDR")
	if err != nil {
		return &envVars{}, err
	}

	USER_SERVICE_DATABASE_NAME, err := utils.GetEnvVar("USER_SERVICE_DATABASE_NAME")
	if err != nil {
		return &envVars{}, err
	}

	DB_CONNECTION_LINK, err := utils.GetEnvVar("DB_CONNECTION_LINK")
	if err != nil {
		return &envVars{}, err
	}

	USER_SERVICE_DATABASE_COLLECTION, err := utils.GetEnvVar("USER_SERVICE_DATABASE_COLLECTION")
	if err != nil {
		return &envVars{}, err
	}

	SECRET_KEY, err := utils.GetEnvVar("SECRET_KEY")
	if err != nil {
		return &envVars{}, err
	}
	if len(SECRET_KEY) < minSecretKeyLength {
		return &envVars{}, fmt.Errorf("SECRET_KEY must be at least %d characters", minSecretKeyLength)
	}

	TOKEN_DURATION_TIME, err := utils.GetEnvVar("TOKEN_DURATION_TIME")
	if err != nil {
		return &envVars{}, err
	}

	return &envVars{
		USER_SERVICE_GRPC_ADDR:           USER_SERVICE_GRPC_ADDR,
		USER_SERVICE_DATABASE_NAME:       USER_SERVICE_DATABASE_NAME,
		DB_CONNECTION_LINK:               DB_CONNECTION_LINK,
		USER_SERVICE_DATABASE_COLLECTION: USER_SERVICE_DATABASE_COLLECTION,
		SECRET_KEY:                       SECRET_KEY,
		TOKEN_DURATION_TIME:              TOKEN_DURATION_TIME,
	}, nil
}
