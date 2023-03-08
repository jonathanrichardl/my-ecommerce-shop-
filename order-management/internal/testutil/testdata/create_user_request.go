package testdata

import (
	"my-ecommerce-shop/order-management/internal/contract"
)

func CreateUserRequest(userId string, userName string, email string, phoneNumber string) contract.CreateUserRequest {
	return contract.CreateUserRequest{
		Id:          userId,
		Username:    userName,
		Password:    "7772231",
		Email:       email,
		PhoneNumber: phoneNumber,
		Role:        "BUYER",
	}
}
