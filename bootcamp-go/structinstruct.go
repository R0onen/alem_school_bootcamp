package bootcamp

import (
	"bootcamp/firststruct"
)

type Cart struct {
	Owner *firststruct.User
	Items []string
}

func CreateMyCart(owner *firststruct.User, items []string) *Cart {
	return &Cart{
		Owner: owner,
		Items: items,
	}
}
