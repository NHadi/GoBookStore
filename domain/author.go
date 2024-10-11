package domain

import "errors"

// Author represents the author of a book
type Author struct {
	Name string
}

// NewAuthor creates a new Author
func NewAuthor(name string) (Author, error) {
	if name == "" {
		return Author{}, errors.New("author name cannot be empty")
	}
	return Author{Name: name}, nil
}
