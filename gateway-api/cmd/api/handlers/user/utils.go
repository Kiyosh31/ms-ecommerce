package user

import (
	"errors"

	userPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/generated/user-service"
)

func ValidateUserPayload(in *userPb.User) []error {
	var errs []error

	if in.GetName() == "" {
		errs = append(errs, errors.New("missing name"))
	}

	if in.GetAddresses() != nil {
		err := ValidateAddress(in.GetAddresses())
		if err != nil {
			errs = append(errs, err)
		}
	}

	if in.GetEmail() == "" {
		errs = append(errs, errors.New("missing email"))
	}

	if in.GetPassword() == "" {
		errs = append(errs, errors.New("missing password"))
	}

	return errs
}

func ValidateAddress(addresses []*userPb.Address) error {
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

func ValidateReactivateUser(payload *userPb.ReactivateUserRequest) []error {
	var errs []error

	if payload.GetEmail() == "" {
		errs = append(errs, errors.New("missing email"))
	}

	if payload.GetPassword() == "" {
		errs = append(errs, errors.New("missing password"))
	}

	return errs
}
