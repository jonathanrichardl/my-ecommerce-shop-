package controller

import (
	"encoding/json"
	"io"
	"my-ecommerce-shop/order-management/internal/contract"
	"my-ecommerce-shop/order-management/internal/service"
	"net/http"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) *contract.GenericResponse {
	reqBody, err := io.ReadAll(r.Body)

	if err != nil {
		return contract.BuildErrorResponse("Unable to parse input")
	}

	var createUserRequest contract.CreateUserRequest
	err = json.Unmarshal(reqBody, &createUserRequest)
	if err != nil {
		return contract.BuildErrorResponse("Unable to parse input")
	}

	response, err := c.userService.Create(createUserRequest)
	if err != nil {
		return contract.BuildErrorResponse("Unable to create user")
	}

	return contract.BuildResponse(response)
}
