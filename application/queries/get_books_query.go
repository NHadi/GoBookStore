package queries

import (
	"BookStore/domain/entities"
	"BookStore/domain/repositories"
)

// GetBooksQuery represents the query to retrieve all books
type GetBooksQuery struct{}

// Handle executes the query to retrieve all books
func (q GetBooksQuery) Handle(repo repositories.BookRepository) []*entities.Book {
	return repo.List()
}
