package repositories

import "BookStore/domain/entities"

// BookRepository defines methods for interacting with books
type BookRepository interface {
	Add(book *entities.Book)
	List() []*entities.Book
	FindById(id string) (*entities.Book, error)
	Update(book *entities.Book) error
	Delete(id string) error
	// following methods for Elasticsearch operations
	IndexBook(book *entities.Book) error
	SearchBooks(title string) ([]entities.Book, error)
}
