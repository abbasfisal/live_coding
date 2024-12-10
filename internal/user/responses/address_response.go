package responses

import (
	"live_coding/entity"
)

type Address struct {
	Street  string `json:"street,omitempty"`
	City    string `json:"city,omitempty"`
	ZIPCode string `json:"zip_code,omitempty"`
	Country string `json:"country,omitempty"`
}

func ToAddresses(addressCollection []entity.Address) []Address {
	var addresses []Address
	for _, address := range addressCollection {
		addresses = append(addresses, toAddress(address))
	}
	return addresses
}

func toAddress(address entity.Address) Address {
	return Address{
		Street:  address.Street,
		City:    address.City,
		ZIPCode: address.ZipCode,
		Country: address.Country,
	}
}
