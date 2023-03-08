package testutil

import (
	"github.com/stretchr/testify/mock"
	"my-ecommerce-shop/order-management/internal/repository/model"
)

type UserRepositoryMock struct {
	mock.Mock
}

func NewUserRepositoryMock() *UserRepositoryMock { return &UserRepositoryMock{} }

func (m *UserRepositoryMock) Get(id string) (*model.User, error) {
	args := m.Called(id)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *UserRepositoryMock) Create(user model.User) (string, error) {
	args := m.Called(user)
	return args.String(0), args.Error(1)
}
