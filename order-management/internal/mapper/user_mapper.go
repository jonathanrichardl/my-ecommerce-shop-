package mapper

import (
	"my-ecommerce-shop/order-management/internal/contract"
	"my-ecommerce-shop/order-management/internal/repository/model"
)

func MapToUserModel(user contract.CreateUserRequest, id string, hashedPassword string) model.User {
	return model.User{
		Id:          id,
		Username:    user.Username,
		Password:    hashedPassword,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Role:        model.GetRole(user.Role),
	}
}
