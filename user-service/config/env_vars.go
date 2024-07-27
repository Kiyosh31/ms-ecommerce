package config

import "github.com/Kiyosh31/ms-ecommerce-common/utils"

type envVars struct {
	USER_SERVICE_GRPC_ADDR string `env:"GATEWAY_API_HTTP_ADRR"`
}

var GlobalEnvVars envVars

func LoadEnvVars() []error {
	var errors []error

	envVarsMap := map[string]*string{
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
