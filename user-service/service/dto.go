package service

import (
	"fmt"

	"github.com/Kiyosh31/ms-ecommerce-common/database"
	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"

	userPb "github.com/Kiyosh31/ms-ecommerce/user-service/proto"
	"github.com/Kiyosh31/ms-ecommerce/user-service/user_types"
)

func createUserPbDto(in user_types.UserSchema) userPb.User {
	return userPb.User{
		Id:       in.ID.Hex(),
		Name:     in.Name,
		Email:    in.Email,
		Password: in.Password,
		// Orders:    in.Orders,
		Addresses: createAddressTypeDto(in.Addresses),
		IsActive:  in.IsActive,
	}
}

func createAddressPbDto(in []*userPb.Address) ([]user_types.Address, error) {
	var addresses []user_types.Address

	var mongoId primitive.ObjectID
	var err error

	for _, address := range in {
		if address.GetId() != "" {
			mongoId, err = database.GetMongoId(address.GetId())
			if err != nil {
				return []user_types.Address{}, err
			}
		} else {
			mongoId = primitive.NewObjectID()
		}

		addresses = append(addresses, user_types.Address{
			ID:      mongoId,
			Name:    address.GetName(),
			Street:  address.GetCity(),
			City:    address.GetCity(),
			State:   address.GetState(),
			Country: address.GetCountry(),
			ZipCode: address.GetZipCode(),
		})
	}

	return addresses, nil
}

func createResponsePbDto(message string, id interface{}, in user_types.UserSchema) (userPb.Response, error) {
	var userId primitive.ObjectID
	var ok bool

	if id != nil {
		userId, ok = id.(primitive.ObjectID)
		if !ok {
			return userPb.Response{}, fmt.Errorf("failed to parse _id to string")
		}
	} else {
		userId = in.ID
	}

	return userPb.Response{
		Message: message,
		User: &userPb.User{
			Id:        userId.Hex(),
			Name:      in.Name,
			Addresses: createAddressTypeDto(in.Addresses),
			Email:     in.Email,
			Password:  in.Password,
			IsActive:  in.IsActive,
		},
	}, nil
}

func createAddressTypeDto(in []user_types.Address) []*userPb.Address {
	var addresses []*userPb.Address

	for _, address := range in {
		addresses = append(addresses, &userPb.Address{
			Id:      address.ID.Hex(),
			Street:  address.Street,
			City:    address.City,
			State:   address.State,
			Country: address.Country,
			Name:    address.Name,
			ZipCode: address.ZipCode,
		})
	}

	return addresses
}

func createUserSchemaDto(in *userPb.User) (user_types.UserSchema, error) {

	addresses, err := createAddressPbDto(in.GetAddresses())
	if err != nil {
		return user_types.UserSchema{}, err
	}

	hashedPassword, err := utils.HashPassword(in.GetPassword())
	if err != nil {
		return user_types.UserSchema{}, err
	}

	var userID primitive.ObjectID
	if in.Id != "" {
		userID, err = database.GetMongoId(in.GetId())
		if err != nil {
			return user_types.UserSchema{}, err
		}
	}

	return user_types.UserSchema{
		ID:        userID,
		Name:      in.GetName(),
		Email:     in.GetEmail(),
		Password:  hashedPassword,
		Addresses: addresses,
		IsActive:  in.GetIsActive(),
	}, nil
}
