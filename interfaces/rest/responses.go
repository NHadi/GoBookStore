package rest

import "BookStore/domain/entities"

// BookResponse represents the response for a single book.
// swagger:model
type BooksResponse struct {
	Count int              `json:"count"`
	Books []*entities.Book `json:"books"`
}

// ErrorBody defines the structure for error messages.
type ErrorBody struct {
	Message string `json:"message"`
}

// ErrorResponse represents a standard error response.
// swagger:response errorResponse
type ErrorResponse struct {
	// in: body
	Body ErrorBody `json:"body"`
}
