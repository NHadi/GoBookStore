package entities

import (
	"BookStore/domain/value_objects"
	"errors"
)

// Book represents the aggregate root for a book entity
type Book struct {
	ID     string               `json:"id"`     // ID of the book
	Title  string               `json:"title"`  // Title of the book
	Author value_objects.Author `json:"author"` // Author of the book
}

// NewBook creates a new book
func NewBook(id, title string, author value_objects.Author) (*Book, error) {
	if title == "" {
		return nil, errors.New("book title cannot be empty")
	}
	return &Book{ID: id, Title: title, Author: author}, nil
}

// ChangeTitle changes the title of a book
func (b *Book) ChangeTitle(newTitle string) error {
	if newTitle == "" {
		return errors.New("book title cannot be empty")
	}
	b.Title = newTitle
	return nil
}
