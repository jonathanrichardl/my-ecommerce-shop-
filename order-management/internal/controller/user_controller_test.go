package controller

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"my-ecommerce-shop/order-management/internal/contract"
	"my-ecommerce-shop/order-management/internal/testutil"
	"my-ecommerce-shop/order-management/internal/testutil/testdata"
	"net/http"
	"net/http/httptest"
	"testing"
)

var userService *testutil.UserServiceMock
var userController *UserController

func TestMain(m *testing.M) {
	setUp()
	m.Run()
}

func setUp() {
	userService = testutil.NewUserServiceMock()
	userController = NewUserController(userService)
}

func TestUserController_CreateUser(t *testing.T) {
	createUserRequest := testdata.CreateUserRequest("userID", "jorich", "jonathanrichard@h.com", "000023")
	requestBody, _ := json.Marshal(createUserRequest)
	expectedResponse := contract.CreateUserResponse{UserId: "userID"}
	req, err := http.NewRequest("POST", "/signup", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	userService.On("Create", createUserRequest).Return(expectedResponse, nil)
	rr := httptest.NewRecorder()

	response := userController.CreateUser(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, response.Data.(contract.CreateUserResponse), expectedResponse)
}
