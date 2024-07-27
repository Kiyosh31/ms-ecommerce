package config

import (
	"github.com/Kiyosh31/ms-ecommerce-common/utils"
)

type envVars struct {
	GATEWAY_API_HTTP_ADRR  string `env:"GATEWAY_API_HTTP_ADRR"`
	USER_SERVICE_GRPC_ADDR string `env:"USER_SERVICE_GRPC_ADDR"`
}

var GlobalEnvVars envVars

func LoadEnvVars() []error {
	var errors []error

	envVarsMap := map[string]*string{
		"GATEWAY_API_HTTP_ADDR":  &GlobalEnvVars.GATEWAY_API_HTTP_ADRR,
		"USER_SERVICE_GRPC_ADDR": &GlobalEnvVars.USER_SERVICE_GRPC_ADDR,
	}

	for key, addr := range envVarsMap {
		val, err := utils.GetEnvVar(key)
		if err != nil {
			errors = append(errors, err)
		} else {
			*addr = val
		}
	}

	return errors
}
