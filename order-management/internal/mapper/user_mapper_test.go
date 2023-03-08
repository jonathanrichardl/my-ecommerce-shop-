package mapper

import (
	"github.com/stretchr/testify/assert"
	"my-ecommerce-shop/order-management/internal/contract"
	"my-ecommerce-shop/order-management/internal/repository/model"
	"my-ecommerce-shop/order-management/internal/util"
	"testing"
)

func TestMapToUserModel(t *testing.T) {
	userId := util.NewUUID()
	createUserRequest := contract.CreateUserRequest{
		Username:    "jorich",
		Password:    "7772231",
		Email:       "jonathan12345@jonathan.com",
		PhoneNumber: "0123443445",
		Role:        "BUYER",
	}
	expectedUserModel := model.User{
		Id:          userId,
		Username:    "jorich",
		Password:    util.EncryptPassword("7772231"),
		Email:       "jonathan12345@jonathan.com",
		PhoneNumber: "0123443445",
		Role:        model.BUYER,
	}

	actualUserModel := MapToUserModel(createUserRequest, userId, util.EncryptPassword("7772231"))

	assert.Equal(t, expectedUserModel, actualUserModel)
}
