package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"my-ecommerce-shop/order-management/internal/testutil"
	"my-ecommerce-shop/order-management/internal/testutil/testdata"
	"testing"
)

var repo *testutil.UserRepositoryMock
var userService UserService

func TestMain(m *testing.M) {
	setUp()
	m.Run()
}

func TestUserService_Create_ShouldCreateSuccesfully(t *testing.T) {
	createUserRequest := testdata.CreateUserRequest("userId1", "jorich1", "jonathan123457@jonathan.com",
		"01234434457")

	repo.On("Create", mock.Anything).Return("returnedid", nil)

	response, err := userService.Create(createUserRequest)

	assert.Equal(t, "returnedid", response.UserId)
	assert.NoError(t, err)
}

func setUp() {
	repo = testutil.NewUserRepositoryMock()
	userService = NewUserService(repo)
}
