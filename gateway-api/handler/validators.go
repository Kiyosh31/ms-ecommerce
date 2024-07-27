package handler

import (
	"errors"

	userPb "github.com/Kiyosh31/ms-ecommerce/user-service/proto"
)

func validateUserPayload(payload *userPb.User) []error {
	var errs []error

	if payload.GetId() == "" {
		errs = append(errs, errors.New("missing ID"))
	}

	if payload.GetName() == "" {
		errs = append(errs, errors.New("missing name"))
	}

	if payload.GetLastName() == "" {
		errs = append(errs, errors.New("missing last name"))
	}

	if payload.Birth == "" {
		errs = append(errs, errors.New("missing birth"))
	}

	if payload.GetCards() == nil {
		errs = append(errs, errors.New("missing card(s)"))
	}

	if payload.GetAddress() == nil {
		errs = append(errs, errors.New("missing address(s)"))
	}

	if payload.GetEmail() == "" {
		errs = append(errs, errors.New("missing email"))
	}

	if payload.GetPassword() == "" {
		errs = append(errs, errors.New("missing password"))
	}

	return errs
}
