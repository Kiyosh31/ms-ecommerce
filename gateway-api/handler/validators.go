package handler

import (
	"errors"

	cartPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/cart-service"
	productPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/product-service"
	userPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/user-service"
)

// user-service
func validateUserPayload(payload *userPb.User) []error {
	var errs []error

	if payload.GetFirstName() == "" {
		errs = append(errs, errors.New("missing first name"))
	}

	if payload.GetLastName() == "" {
		errs = append(errs, errors.New("missing last name"))
	}

	if payload.GetBirthDate() == "" {
		errs = append(errs, errors.New("missing birth date"))
	}

	if payload.GetCards() != nil {
		err := validateCards(payload.GetCards())
		if err != nil {
			errs = append(errs, err)
		}
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

func validateCards(cards []*userPb.Card) error {
	for _, card := range cards {
		if card.GetNumber() == "" {
			return errors.New("card number is required")
		}

		if card.GetCardHolderName() == "" {
			return errors.New("card holder name is required")
		}

		if card.GetCvv() == 0 {
			return errors.New("card cvv is required")
		}

		if card.GetExpiration() == "" {
			return errors.New("card expiration is required")
		}

		if card.GetCardType() == "" {
			return errors.New("card type is required")
		}

		if !card.GetIsDefault() {
			return errors.New("card default is required")
		}
	}

	return nil
}

func validateAddress(addresses []*userPb.Address) error {
	for _, address := range addresses {
		if address.GetName() == "" {
			return errors.New("address must contain a name")
		}

		if address.GetZipCode() == 0 {
			return errors.New("address must contain a zip code")
		}

		if !address.GetIsDefault() {
			return errors.New("address must contain a default")
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

	return errs
}

// cart-service
func validateCartPayload(payload *cartPb.Cart) []error {
	var errs []error

	if payload.GetId() == "" {
		errs = append(errs, errors.New("missing id"))
	}

	if payload.GetUserId() == "" {
		errs = append(errs, errors.New("missing userId"))
	}

	if payload.GetId() == "" {
		errs = append(errs, errors.New("missing id"))
	}

	if payload.GetTotal() == 0 {
		errs = append(errs, errors.New("missing total"))
	}

	return errs
}
