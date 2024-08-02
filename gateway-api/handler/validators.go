package handler

import (
	"errors"

	productPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/product-service"
	userPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/user-service"
)

// user-service
func validateUserPayload(payload *userPb.User) []error {
	var errs []error

	if payload.GetName() == "" {
		errs = append(errs, errors.New("missing name"))
	}

	if payload.GetAddresses() != nil {
		err := validateAddress(payload.GetAddresses())
		if err != nil {
			errs = append(errs, err)
		}
	}

	if payload.GetEmail() == "" {
		errs = append(errs, errors.New("missing email"))
	}

	if payload.GetPassword() == "" {
		errs = append(errs, errors.New("missing password"))
	}

	return errs
}

func validateAddress(addresses []*userPb.Address) error {
	for _, address := range addresses {
		if address.GetName() == "" {
			return errors.New("address must contain a name")
		}

		if address.GetCity() == "" {
			return errors.New("address must contain a city")
		}

		if address.GetCountry() == "" {
			return errors.New("address must contain a country")
		}

		if address.GetState() == "" {
			return errors.New("address must contain a state")
		}

		if address.GetStreet() == "" {
			return errors.New("address must contain a street")
		}

		if address.GetZipCode() == 0 {
			return errors.New("address must contain a zip code")
		}
	}

	return nil
}

func validateReactivateUser(payload *userPb.ReactivarUserRequest) []error {
	var errs []error

	if payload.GetEmail() == "" {
		errs = append(errs, errors.New("missing email"))
	}

	if payload.GetPassword() == "" {
		errs = append(errs, errors.New("missing password"))
	}

	return errs
}

// product-service
func validateProductPayload(payload *productPb.Product) []error {
	var errs []error

	if payload.GetSellerId() == "" {
		errs = append(errs, errors.New("missing sellerId"))
	}

	if payload.GetName() == "" {
		errs = append(errs, errors.New("missing name"))
	}

	if payload.GetDescription() == "" {
		errs = append(errs, errors.New("missing description"))
	}

	if payload.GetPrice() == 0 {
		errs = append(errs, errors.New("missing price"))
	}

	if payload.GetAvailableQuantity() == 0 {
		errs = append(errs, errors.New("missing availableQuantity"))
	}

	return errs
}
