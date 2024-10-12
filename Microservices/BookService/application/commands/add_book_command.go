package commands

import (
	"BookStore/domain/entities"
	"BookStore/domain/repositories"
	"BookStore/domain/value_objects"
)

// AddBookCommand represents the command to add a book
type AddBookCommand struct {
	ID     string
	Title  string
	Author string
}

// Handle executes the command to add a book
func (cmd AddBookCommand) Handle(repo repositories.BookRepository) (*entities.Book, error) {
	author, err := value_objects.NewAuthor(cmd.Author)
	if err != nil {
		return nil, err
	}

	book, err := entities.NewBook(cmd.ID, cmd.Title, author)
	if err != nil {
		return nil, err
	}

	repo.Add(book)
	return book, nil
}
