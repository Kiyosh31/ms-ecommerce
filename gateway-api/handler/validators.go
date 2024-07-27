package handler

import (
	"errors"

	userPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/user-service"
)

func validateUserPayload(payload *userPb.User) []error {
	var errs []error

	if payload.GetName() == "" {
		errs = append(errs, errors.New("missing name"))
	}

	if payload.GetLastName() == "" {
		errs = append(errs, errors.New("missing last name"))
	}

	if payload.GetBirth() == "" {
		errs = append(errs, errors.New("missing birth"))
	}

	if payload.GetCards() != nil {
		errs = append(errs, validateCards(payload.GetCards()))
	}

	if payload.GetAddresses() != nil {
		errs = append(errs, validateAddress(payload.GetAddresses()))
	}

	if payload.GetEmail() == "" {
		errs = append(errs, errors.New("missing email"))
	}

	if payload.GetPassword() == "" {
		errs = append(errs, errors.New("missing password"))
	}

	return errs
}

func validateCards(cards *userPb.CardList) error {
	for _, card := range cards.Cards {
		if card.GetNumber() == 0 {
			return errors.New("card number is required")
		}

		if card.GetCvv() == 0 {
			return errors.New("card cvv is required")
		}

		if card.GetExpiration() == "" {
			return errors.New("card expiration is required")
		}

		if !card.GetDefault() {
			return errors.New("card default is required")
		}
	}

	return nil
}

func validateAddress(addresses *userPb.AddressList) error {
	for _, address := range addresses.Address {
		if address.GetName() == "" {
			return errors.New("address must contain a name")
		}

		if address.GetZipCode() == 0 {
			return errors.New("address must contain a zip code")
		}

		if !address.GetDefault() {
			return errors.New("address must contain a default")
		}
	}
	return nil
}
