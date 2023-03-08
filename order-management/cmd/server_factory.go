package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"my-ecommerce-shop/order-management/internal/controller"
	"my-ecommerce-shop/order-management/internal/repository"
	"my-ecommerce-shop/order-management/internal/service"
	"os"
)

func newDb(databaseUrl string) *sql.DB {
	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		panic(err)
	}

	return db
}

func newRouter(userController *controller.UserController) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/signup", controller.Wrap(userController.CreateUser)).Methods("POST")

	return router
}

func createServer() *mux.Router {
	godotenv.Load()

	db := newDb(os.Getenv("DB_HOST"))

	userRepository := repository.NewUserRepository(db)

	userService := service.NewUserService(userRepository)

	userController := controller.NewUserController(userService)

	router := newRouter(userController)

	return router
}
