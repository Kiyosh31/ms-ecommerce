package handler

import (
	"errors"

	inventoryPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/inventory-service"
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
		categoryErrors := validateCategoryPayload(payload.GetCategory())
		errs = append(errs, categoryErrors...)
	}

	if payload.GetBrand() == nil {
		errs = append(errs, errors.New("missing brand"))
	} else {
		brandErrors := validateBrandPayload(payload.GetBrand())
		errs = append(errs, brandErrors...)
	}

	if payload.GetImages() == nil {
		errs = append(errs, errors.New("missing images"))
	}

	if payload.GetAttributes() == nil {
		errs = append(errs, errors.New("missing attributes"))
	} else {
		attrErrors := validateAttributesPayload(payload.GetAttributes())
		errs = append(errs, attrErrors...)
	}

	if payload.GetInventory() == nil {
		errs = append(errs, errors.New("missing inventory"))
	}

	return errs
}

// category
func validateCategoryPayload(payload *productPb.Category) []error {
	var errs []error

	if payload.GetName() == "" {
		errs = append(errs, errors.New("category: Missing name"))
	}

	if payload.GetDescription() == "" {
		errs = append(errs, errors.New("category: Missing description"))
	}

	return errs
}

// brand
func validateBrandPayload(payload *productPb.Brand) []error {
	var errs []error

	if payload.GetName() == "" {
		errs = append(errs, errors.New("brand: Missing name"))
	}

	if payload.GetDescription() == "" {
		errs = append(errs, errors.New("brand: Missing description"))
	}

	return errs
}

// attributes
func validateAttributesPayload(payload *productPb.Attributes) []error {
	var errs []error

	if payload.GetSize() == 0 {
		errs = append(errs, errors.New("attributes: missing size"))
	}

	if payload.GetColor() == "" {
		errs = append(errs, errors.New("attributes: missing color"))
	}

	if payload.GetMaterial() == "" {
		errs = append(errs, errors.New("attributes: missing material"))
	}

	return errs
}

// inventory service
func validateInventoryPayload(payload *inventoryPb.Inventory) []error {
	var errs []error

	if payload.GetProductId() == "" {
		errs = append(errs, errors.New("missing productId"))
	}

	if payload.GetLocation() == "" {
		errs = append(errs, errors.New("missing location"))
	}

	if payload.GetQuantity() == 0 {
		errs = append(errs, errors.New("missing quantity"))
	}

	return errs
}
