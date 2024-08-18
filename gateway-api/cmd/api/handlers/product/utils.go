package product

import (
	"errors"

	productPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/cmd/generated/product-service"
)

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
		categoryErrors := ValidateCategoryPayload(payload.GetCategory())
		errs = append(errs, categoryErrors...)
	}

	if payload.GetBrand() == nil {
		errs = append(errs, errors.New("missing brand"))
	} else {
		brandErrors := ValidateBrandPayload(payload.GetBrand())
		errs = append(errs, brandErrors...)
	}

	if payload.GetImages() == nil {
		errs = append(errs, errors.New("missing images"))
	}

	if payload.GetAttributes() == nil {
		errs = append(errs, errors.New("missing attributes"))
	} else {
		attrErrors := ValidateAttributesPayload(payload.GetAttributes())
		errs = append(errs, attrErrors...)
	}

	if payload.GetInventory() == nil {
		errs = append(errs, errors.New("missing inventory"))
	}

	return errs
}

// category
func ValidateCategoryPayload(payload *productPb.Category) []error {
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
func ValidateBrandPayload(payload *productPb.Brand) []error {
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
func ValidateAttributesPayload(payload *productPb.Attributes) []error {
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
