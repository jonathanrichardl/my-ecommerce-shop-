package service

import (
	"my-ecommerce-shop/order-management/internal/contract"
	"my-ecommerce-shop/order-management/internal/mapper"
	"my-ecommerce-shop/order-management/internal/repository"
	"my-ecommerce-shop/order-management/internal/util"
)

type UserService interface {
	Create(request contract.CreateUserRequest) (contract.CreateUserResponse, error)
}

type UserServiceImplementation struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *UserServiceImplementation {
	return &UserServiceImplementation{repository: repository}
}

func (s *UserServiceImplementation) Create(request contract.CreateUserRequest) (contract.CreateUserResponse, error) {
	hashedPassword := util.EncryptPassword(request.Password)
	id := util.NewUUID()
	userToPersist := mapper.MapToUserModel(request, id, hashedPassword)

	id, err := s.repository.Create(userToPersist)

	return contract.CreateUserResponse{UserId: id}, err
}
