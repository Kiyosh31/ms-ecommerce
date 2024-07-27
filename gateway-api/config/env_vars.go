package config

import "github.com/Kiyosh31/ms-ecommerce-common/utils"

type envVars struct {
	GATEWAY_API_HTTP_ADRR  string
	USER_SERVICE_GRPC_ADDR string
}

func LoadEnvVars() (envVars, error) {
	GATEWAY_API_HTTP_ADRR, err := utils.GetEnvVar("GATEWAY_API_HTTP_ADRR")
	if err != nil {
		return envVars{}, err
	}
	// GATEWAY_API_HTTP_ADRR := "localhost:3000"

	USER_SERVICE_GRPC_ADDR, err := utils.GetEnvVar("USER_SERVICE_GRPC_ADDR")
	if err != nil {
		return envVars{}, err
	}
	// USER_SERVICE_GRPC_ADDR := "localhost:3001"

	return envVars{
		GATEWAY_API_HTTP_ADRR:  GATEWAY_API_HTTP_ADRR,
		USER_SERVICE_GRPC_ADDR: USER_SERVICE_GRPC_ADDR,
	}, nil
}
