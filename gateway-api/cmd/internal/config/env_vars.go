package config

import (
	"fmt"

	"github.com/Kiyosh31/ms-ecommerce-common/utils"
)

type envVars struct {
	GATEWAY_API_HTTP_ADRR                   string
	GATEWAY_API_USER_SERVICE_GRPC_ADDR      string
	GATEWAY_API_PRODUCT_SERVICE_GRPC_ADDR   string
	GATEWAY_API_INVENTORY_SERVICE_GRPC_ADDR string
	GATEWAY_API_ORDER_SERVICE_GRPC_ADDR     string
	GATEWAY_API_PAYMENT_SERVICE_GRPC_ADDR   string
	SECRET_KEY                              string
}

const minSecretKeyLength = 32

func LoadEnvVars() (*envVars, error) {
	GATEWAY_API_HTTP_ADRR, err := utils.GetEnvVar("GATEWAY_API_HTTP_ADRR")
	if err != nil {
		return &envVars{}, err
	}

	GATEWAY_API_USER_SERVICE_GRPC_ADDR, err := utils.GetEnvVar("GATEWAY_API_USER_SERVICE_GRPC_ADDR")
	if err != nil {
		return &envVars{}, err
	}

	GATEWAY_API_PRODUCT_SERVICE_GRPC_ADDR, err := utils.GetEnvVar("GATEWAY_API_PRODUCT_SERVICE_GRPC_ADDR")
	if err != nil {
		return &envVars{}, err
	}

	GATEWAY_API_INVENTORY_SERVICE_GRPC_ADDR, err := utils.GetEnvVar("GATEWAY_API_INVENTORY_SERVICE_GRPC_ADDR")
	if err != nil {
		return &envVars{}, err
	}

	GATEWAY_API_ORDER_SERVICE_GRPC_ADDR, err := utils.GetEnvVar("GATEWAY_API_ORDER_SERVICE_GRPC_ADDR")
	if err != nil {
		return &envVars{}, err
	}

	GATEWAY_API_PAYMENT_SERVICE_GRPC_ADDR, err := utils.GetEnvVar("GATEWAY_API_PAYMENT_SERVICE_GRPC_ADDR")
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

	return &envVars{
		GATEWAY_API_HTTP_ADRR:                   GATEWAY_API_HTTP_ADRR,
		GATEWAY_API_USER_SERVICE_GRPC_ADDR:      GATEWAY_API_USER_SERVICE_GRPC_ADDR,
		GATEWAY_API_PRODUCT_SERVICE_GRPC_ADDR:   GATEWAY_API_PRODUCT_SERVICE_GRPC_ADDR,
		GATEWAY_API_INVENTORY_SERVICE_GRPC_ADDR: GATEWAY_API_INVENTORY_SERVICE_GRPC_ADDR,
		GATEWAY_API_ORDER_SERVICE_GRPC_ADDR:     GATEWAY_API_ORDER_SERVICE_GRPC_ADDR,
		GATEWAY_API_PAYMENT_SERVICE_GRPC_ADDR:   GATEWAY_API_PAYMENT_SERVICE_GRPC_ADDR,
	}, nil
}
