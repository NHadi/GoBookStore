package value_objects

import "errors"

// Author represents the value object for a book's author
type Author struct {
	Name string `json:"name"` // The author's name
}

// NewAuthor creates a new Author value object
func NewAuthor(name string) (Author, error) {
	if name == "" {
		return Author{}, errors.New("author name cannot be empty")
	}
	return Author{Name: name}, nil
}
