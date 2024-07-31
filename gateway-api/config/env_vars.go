package config

import "github.com/Kiyosh31/ms-ecommerce-common/utils"

type envVars struct {
	GATEWAY_API_HTTP_ADRR                 string
	GATEWAY_API_USER_SERVICE_GRPC_ADDR    string
	GATEWAY_API_PRODUCT_SERVICE_GRPC_ADDR string
	GATEWAY_API_CART_SERVICE_GRPC_ADDR    string
}

func LoadEnvVars() (envVars, error) {
	GATEWAY_API_HTTP_ADRR, err := utils.GetEnvVar("GATEWAY_API_HTTP_ADRR")
	if err != nil {
		return envVars{}, err
	}

	GATEWAY_API_USER_SERVICE_GRPC_ADDR, err := utils.GetEnvVar("GATEWAY_API_USER_SERVICE_GRPC_ADDR")
	if err != nil {
		return envVars{}, err
	}

	GATEWAY_API_PRODUCT_SERVICE_GRPC_ADDR, err := utils.GetEnvVar("GATEWAY_API_PRODUCT_SERVICE_GRPC_ADDR")
	if err != nil {
		return envVars{}, err
	}

	GATEWAY_API_CART_SERVICE_GRPC_ADDR, err := utils.GetEnvVar("GATEWAY_API_CART_SERVICE_GRPC_ADDR")
	if err != nil {
		return envVars{}, err
	}

	return envVars{
		GATEWAY_API_HTTP_ADRR:                 GATEWAY_API_HTTP_ADRR,
		GATEWAY_API_USER_SERVICE_GRPC_ADDR:    GATEWAY_API_USER_SERVICE_GRPC_ADDR,
		GATEWAY_API_PRODUCT_SERVICE_GRPC_ADDR: GATEWAY_API_PRODUCT_SERVICE_GRPC_ADDR,
		GATEWAY_API_CART_SERVICE_GRPC_ADDR:    GATEWAY_API_CART_SERVICE_GRPC_ADDR,
	}, nil
}
