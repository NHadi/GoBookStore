package repositories

import (
	"BookStore/domain/entities"
	"BookStore/domain/repositories"
	"errors"
)

// InMemoryBookRepository is an in-memory implementation of BookRepository
type InMemoryBookRepository struct {
	books map[string]*entities.Book
}

func NewInMemoryBookRepository() repositories.BookRepository {
	return &InMemoryBookRepository{
		books: make(map[string]*entities.Book),
	}
}

func (r *InMemoryBookRepository) Add(book *entities.Book) {
	r.books[book.ID] = book
}

func (r *InMemoryBookRepository) List() []*entities.Book {
	var result []*entities.Book
	for _, book := range r.books {
		result = append(result, book)
	}
	return result
}

func (r *InMemoryBookRepository) FindById(id string) (*entities.Book, error) {
	if book, exists := r.books[id]; exists {
		return book, nil
	}
	return nil, errors.New("book not found")
}

func (r *InMemoryBookRepository) Update(book *entities.Book) error {
	if _, exists := r.books[book.ID]; !exists {
		return errors.New("book not found")
	}
	r.books[book.ID] = book
	return nil
}

func (r *InMemoryBookRepository) Delete(id string) error {
	if _, exists := r.books[id]; !exists {
		return errors.New("book not found")
	}
	delete(r.books, id)
	return nil
}
