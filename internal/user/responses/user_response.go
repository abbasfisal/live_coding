package responses

import (
	"github.com/google/uuid"
	"live_coding/internal/user/entity"

	"strings"
)

type User struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Family    string    `json:"family,omitempty"`
	Addresses []Address `json:"addresses,omitempty"`
}

func ToUser(user entity.User) User {

	name, family := getNameAndFamily(user.Name)
	return User{
		ID:        user.ID,
		Name:      name,
		Family:    family,
		Addresses: ToAddresses(user.Addresses),
	}
}

func getNameAndFamily(fullName string) (name string, family string) {
	parts := strings.Split(fullName, " ")
	if len(parts) > 1 {
		return parts[0], parts[1]
	}
	return "", ""
}
