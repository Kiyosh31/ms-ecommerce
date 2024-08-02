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

func validateReactivateUser(payload *userPb.ReactivateUserRequest) []error {
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

	if payload.GetName() == "" {
		errs = append(errs, errors.New("missing name"))
	}

	if payload.GetDescription() == "" {
		errs = append(errs, errors.New("missing description"))
	}

	if payload.GetPrice() == 0 {
		errs = append(errs, errors.New("missing price"))
	}

	if payload.GetCategory() == nil {
		errs = append(errs, errors.New("missing category"))
	} else {
		if payload.GetCategory().GetId() == "" {
			errs = append(errs, errors.New("category: Missing category id"))
		}
		if payload.GetCategory().GetName() == "" {
			errs = append(errs, errors.New("category: Missing name"))
		}

		if payload.GetCategory().GetDescription() == "" {
			errs = append(errs, errors.New("category: Missing description"))
		}

		if payload.GetCategory().GetParentCategory() == "" {
			errs = append(errs, errors.New("category: Missing parentCategory"))
		}
	}

	if payload.GetBrand() == nil {
		errs = append(errs, errors.New("missing brand"))
	} else {
		if payload.GetBrand().GetId() == "" {
			errs = append(errs, errors.New("brand: Missing brand id"))
		}

		if payload.GetBrand().GetName() == "" {
			errs = append(errs, errors.New("brand: Missing name"))
		}

		if payload.GetBrand().GetDescription() == "" {
			errs = append(errs, errors.New("brand: Missing description"))
		}
	}

	if payload.GetImages() == nil {
		errs = append(errs, errors.New("missing images"))
	}

	if payload.GetAttributes() == nil {
		errs = append(errs, errors.New("missing attributes"))
	} else {
		if payload.GetAttributes().GetSize() == 0 {
			errs = append(errs, errors.New("attributes: missing size"))
		}

		if payload.GetAttributes().GetColor() == "" {
			errs = append(errs, errors.New("attributes: missing color"))
		}

		if payload.GetAttributes().GetMaterial() == "" {
			errs = append(errs, errors.New("attributes: missing material"))
		}
	}

	if payload.GetInventory() == nil {
		errs = append(errs, errors.New("missing inventory"))
	}

	return errs
}
