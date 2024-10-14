// main.go

package main

import (
	"log"
	"net/http"

	"BookStore/application/services"
	"BookStore/infrastructure/repositories"
	"BookStore/interfaces/rest"

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
	// Define the UserService URL
	gatewayServiceURL := "http://localhost:8080"

	// Initialize Cassandra repository
	cassandraRepo, err := repositories.NewCassandraBookRepository()
	if err != nil {
		log.Fatalf("Failed to initialize Cassandra repository: %v", err)
	}

	// Initialize Elasticsearch repository
	elasticRepo, err := repositories.NewElasticsearchBookRepository()
	if err != nil {
		log.Fatalf("Failed to initialize Elasticsearch repository: %v", err)
	}

	// Initialize your application service with both repositories
	bookService := services.NewBookApplicationService(cassandraRepo, elasticRepo, gatewayServiceURL)

	// Initialize your REST controller
	bookController := rest.NewBookController(bookService)

	// Set up the router
	r := mux.NewRouter()

	// Define your routes
	r.HandleFunc("/books", bookController.CreateBook).Methods("POST")
	r.HandleFunc("/books", bookController.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", bookController.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", bookController.DeleteBook).Methods("DELETE")
	r.HandleFunc("/books/borrow/{userId}/{bookId}", bookController.BorrowBook).Methods("POST")
	r.HandleFunc("/books/search", bookController.SearchBooks).Methods("GET")
	// Swagger route
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Health check endpoint
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Healthy"))
	})

	// Start the server
	log.Println("Starting server on :8082")
	if err := http.ListenAndServe(":8082", r); err != nil {
		log.Fatal(err)
	}
}
