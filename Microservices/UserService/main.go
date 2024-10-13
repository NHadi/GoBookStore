package main

import (
	"UserService/application/services"
	_ "UserService/docs"
	"UserService/infrastructure/repositories"
	"UserService/interfaces/rest"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title UserService API
// @version 1.0
// @description UserService for managing user accounts

// @host localhost:8081
// @BasePath /

func main() {
	// Initialize the repository
	userRepo := repositories.NewInMemoryUserRepository()

	// Initialize the application service
	userService := services.NewUserApplicationService(userRepo)

	// Initialize the REST controller
	userController := rest.NewUserController(userService)

	// Set up the router
	r := mux.NewRouter()

	// Define the routes
	r.HandleFunc("/users", userController.GetAllUsers).Methods("GET")
	r.HandleFunc("/users", userController.RegisterUser).Methods("POST")
	r.HandleFunc("/users/{id}", userController.GetUser).Methods("GET")
	r.HandleFunc("/users/{userId}/eligibility", userController.CheckUserEligibility).Methods("GET")

	// Swagger route
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Start the server
	log.Println("Starting UserService on :8081")
	if err := http.ListenAndServe(":8081", r); err != nil {
		log.Fatal(err)
	}
}
