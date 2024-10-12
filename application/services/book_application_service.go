package services

import (
	"BookStore/application/commands"
	"BookStore/application/queries"
	"BookStore/domain/entities"
	"BookStore/domain/repositories"
)

type BookApplicationService struct {
	repo repositories.BookRepository
}

func NewBookApplicationService(repo repositories.BookRepository) *BookApplicationService {
	return &BookApplicationService{repo: repo}
}

func (s *BookApplicationService) AddBook(cmd commands.AddBookCommand) (*entities.Book, error) {
	return cmd.Handle(s.repo)
}

func (s *BookApplicationService) ListBooks(q queries.GetBooksQuery) []*entities.Book {
	return q.Handle(s.repo)
}

// FindByID retrieves a book by its ID.
func (s *BookApplicationService) FindByID(id string) (*entities.Book, error) {
	return s.repo.FindById(id)
}

// Update updates a book's details.
func (s *BookApplicationService) Update(book *entities.Book) error {
	return s.repo.Update(book)
}

// Delete removes a book by its ID.
func (s *BookApplicationService) Delete(id string) error {
	return s.repo.Delete(id)
}
