package userutils

import (
	"github.com/Kiyosh31/ms-ecommerce-common/database"
	"github.com/Kiyosh31/ms-ecommerce/user-service/cmd/internal/domain"
	userPb "github.com/Kiyosh31/ms-ecommerce/user-service/cmd/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MapAddressToProto(in []domain.Address) []*userPb.Address {
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

func MapAddressToDomain(in []*userPb.Address) ([]domain.Address, error) {
	var addresses []domain.Address

	var mongoId primitive.ObjectID
	var err error

	for _, address := range in {
		if address.GetId() != "" {
			mongoId, err = database.GetMongoId(address.GetId())
			if err != nil {
				return []domain.Address{}, err
			}
		} else {
			mongoId = primitive.NewObjectID()
		}

		addresses = append(addresses, domain.Address{
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
