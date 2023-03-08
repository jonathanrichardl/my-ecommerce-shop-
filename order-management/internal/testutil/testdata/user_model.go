package testdata

import (
	"my-ecommerce-shop/order-management/internal/repository/model"
	"my-ecommerce-shop/order-management/internal/util"
)

func CreateTestUser(userId string, userName string, email string, phoneNumber string) model.User {
	return model.User{
		Id:          userId,
		Username:    userName,
		Password:    util.EncryptPassword("7772231"),
		Email:       email,
		PhoneNumber: phoneNumber,
		Role:        model.BUYER,
	}
}
