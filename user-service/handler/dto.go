package handler

import (
	"fmt"

	"github.com/Kiyosh31/ms-ecommerce-common/database"
	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"

	userPb "github.com/Kiyosh31/ms-ecommerce/user-service/proto"
	"github.com/Kiyosh31/ms-ecommerce/user-service/types"
)

func mapUserTypeFromPb(in *userPb.User) (types.UserSchema, error) {
	cards, err := mapCardTypeFromPb(in.GetCards())
	if err != nil {
		return types.UserSchema{}, err
	}

	addresses, err := mapAdressTypeFromPb(in.GetAddresses())
	if err != nil {
		return types.UserSchema{}, err
	}

	hashedPassword, err := utils.HashPassword(in.GetPassword())
	if err != nil {
		return types.UserSchema{}, err
	}

	return types.UserSchema{
		FirstName: in.GetFirstName(),
		LastName:  in.GetLastName(),
		BirthDate: in.GetBirthDate(),
		Email:     in.GetEmail(),
		Password:  hashedPassword,
		Cards:     cards,
		Addresses: addresses,
	}, nil
}

func mapCardTypeFromPb(in []*userPb.Card) ([]types.Card, error) {
	var cards []types.Card

	var mongoId primitive.ObjectID
	var err error

	for _, card := range in {
		if card.GetId() != "" {
			mongoId, err = database.GetMongoId(card.GetId())
			if err != nil {
				return []types.Card{}, err
			}
		} else {
			mongoId = primitive.NewObjectID()
		}

		cards = append(cards, types.Card{
			ID:             mongoId,
			Number:         card.GetNumber(),
			CardHolderName: card.GetCardHolderName(),
			CardType:       card.GetCardType(),
			Cvv:            card.GetCvv(),
			Expiration:     card.GetExpiration(),
			IsDefault:      card.GetIsDefault(),
		})
	}

	return cards, nil
}

func mapAdressTypeFromPb(in []*userPb.Address) ([]types.Address, error) {
	var addresses []types.Address

	var mongoId primitive.ObjectID
	var err error

	for _, address := range in {
		if address.GetId() != "" {
			mongoId, err = database.GetMongoId(address.GetId())
			if err != nil {
				return []types.Address{}, err
			}
		} else {
			mongoId = primitive.NewObjectID()
		}

		addresses = append(addresses, types.Address{
			ID:        mongoId,
			Name:      address.GetName(),
			Street:    address.GetCity(),
			City:      address.GetCity(),
			State:     address.GetState(),
			Country:   address.GetCountry(),
			ZipCode:   address.GetZipCode(),
			IsDefault: address.GetIsDefault(),
		})
	}

	return addresses, nil
}

func mapResponseFromType(message string, id interface{}, in types.UserSchema) (userPb.Response, error) {
	userId, ok := id.(primitive.ObjectID)
	if !ok {
		return userPb.Response{}, fmt.Errorf("failed to parse _id to string")
	}

	return userPb.Response{
		Message: message,
		User: &userPb.User{
			Id:        userId.Hex(),
			FirstName: in.FirstName,
			LastName:  in.LastName,
			BirthDate: in.BirthDate,
			Cards:     mapCardPbToType(in.Cards),
			Addresses: mapAddressPbToType(in.Addresses),
			Email:     in.Email,
			Password:  in.Password,
		},
	}, nil
}

func mapCardPbToType(in []types.Card) []*userPb.Card {
	var cards []*userPb.Card

	for _, card := range in {
		cards = append(cards, &userPb.Card{
			Id:         card.ID.Hex(),
			Number:     card.Number,
			Cvv:        int32(card.Cvv),
			Expiration: card.Expiration,
			IsDefault:  card.IsDefault,
		})
	}

	return cards
}

func mapAddressPbToType(in []types.Address) []*userPb.Address {
	var addresses []*userPb.Address

	for _, address := range in {
		addresses = append(addresses, &userPb.Address{
			Id:        address.ID.Hex(),
			Name:      address.Name,
			ZipCode:   address.ZipCode,
			IsDefault: address.IsDefault,
		})
	}

	return addresses
}
