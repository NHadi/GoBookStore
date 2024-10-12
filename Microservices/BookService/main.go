// main.go

package main

import (
	"log"
	"net/http"

	"BookStore/application/services"
	"BookStore/infrastructure/repositories"
	"BookStore/interfaces/rest"

	_ "BookStore/docs" // Import generated docs

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title BookStore API
// @version 1.0
// @description This is a sample BookStore
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8082
// @BasePath /
func main() {
	// Initialize your repository
	bookRepo := repositories.NewInMemoryBookRepository()

	// Initialize your application service
	bookService := services.NewBookApplicationService(bookRepo)

	// Initialize your REST controller
	bookController := rest.NewBookController(bookService)

	// Set up the router
	r := mux.NewRouter()

	// Define your routes
	r.HandleFunc("/books", bookController.CreateBook).Methods("POST")
	r.HandleFunc("/books", bookController.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", bookController.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", bookController.DeleteBook).Methods("DELETE")

	// Swagger route
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Start the server
	log.Println("Starting server on :8082")
	if err := http.ListenAndServe(":8082", r); err != nil {
		log.Fatal(err)
	}
}
