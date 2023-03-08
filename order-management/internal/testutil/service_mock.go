package testutil

import (
	"github.com/stretchr/testify/mock"
	"my-ecommerce-shop/order-management/internal/contract"
)

type UserServiceMock struct {
	mock.Mock
}

func NewUserServiceMock() *UserServiceMock {
	return &UserServiceMock{}
}

func (m *UserServiceMock) Create(request contract.CreateUserRequest) (contract.CreateUserResponse, error) {
	args := m.Called(request)
	return args.Get(0).(contract.CreateUserResponse), args.Error(1)
}
