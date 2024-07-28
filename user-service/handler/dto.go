package handler

import (
	"github.com/Kiyosh31/ms-ecommerce-common/database"
	"github.com/Kiyosh31/ms-ecommerce-common/utils"

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
	in.Password = hashedPassword

	return types.UserSchema{
		Name:     in.GetName(),
		LastName: in.GetLastName(),
		Birth:    in.GetBirth(),
		Email:    in.GetEmail(),
		Password: in.GetPassword(),
		Card:     cards,
		Address:  addresses,
	}, nil
}

func mapCardTypeFromPb(in *userPb.CardList) ([]types.Card, error) {
	var cards []types.Card

	for _, card := range in.GetCards() {
		mongoId, err := database.GetMongoId(card.GetId())
		if err != nil {
			return []types.Card{}, err
		}

		cards = append(cards, types.Card{
			ID:         mongoId,
			Number:     card.GetNumber(),
			Cvv:        int64(card.GetCvv()),
			Expiration: card.GetExpiration(),
			Default:    card.GetDefault(),
		})
	}

	return cards, nil
}

func mapAdressTypeFromPb(in *userPb.AddressList) ([]types.Address, error) {
	var addresses []types.Address

	for _, address := range in.GetAddress() {
		mongoId, err := database.GetMongoId(address.GetId())
		if err != nil {
			return []types.Address{}, err
		}

		addresses = append(addresses, types.Address{
			ID:      mongoId,
			Name:    address.GetName(),
			ZipCode: address.GetZipCode(),
			Default: address.GetDefault(),
		})
	}

	return addresses, nil
}

func mapResponseFromType(message string, id interface{}, in types.UserSchema) (userPb.Response, error) {
	userId, _ := utils.ParseInterfaceToString(id)
	// if err != nil {
	// 	return userPb.Response{}, err
	// }

	return userPb.Response{
		Message: message,
		User: &userPb.User{
			Id:       userId,
			Name:     in.Name,
			LastName: in.LastName,
			Birth:    in.Birth,
			Cards: &userPb.CardList{
				Cards: mapCardPbToType(in.Card),
			},
			Addresses: &userPb.AddressList{
				Address: mapAddressPbToType(in.Address),
			},
			Email:    in.Email,
			Password: in.Password,
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
			Default:    card.Default,
		})
	}

	return cards
}

func mapAddressPbToType(in []types.Address) []*userPb.Address {
	var addresses []*userPb.Address

	for _, address := range in {
		addresses = append(addresses, &userPb.Address{
			Id:      address.ID.Hex(),
			Name:    address.Name,
			ZipCode: address.ZipCode,
			Default: address.Default,
		})
	}

	return addresses
}
