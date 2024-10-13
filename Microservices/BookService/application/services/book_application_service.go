package services

import (
	"BookStore/application/commands"
	"BookStore/application/queries"
	"BookStore/domain/entities"
	"BookStore/domain/repositories"
	"encoding/json"
	"fmt"
	"net/http"
)

type BookApplicationService struct {
	repo              repositories.BookRepository
	gatewayServiceURL string
}

func NewBookApplicationService(repo repositories.BookRepository, gatewayServiceURL string) *BookApplicationService {
	return &BookApplicationService{repo: repo, gatewayServiceURL: gatewayServiceURL}
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

// BorrowBook checks user eligibility and processes book borrowing
func (s *BookApplicationService) BorrowBook(userId string, bookId string) (string, error) {
	// Check user eligibility by calling UserService
	isEligible, err := s.checkUserEligibility(userId)
	if err != nil {
		return "", err
	}

	if !isEligible {
		return "", fmt.Errorf("user %s is not eligible to borrow books", userId)
	}

	// TODO: Implement logic to mark the book as borrowed (update database, etc.)

	return fmt.Sprintf("User %s successfully borrowed book %s", userId, bookId), nil
}

// checkUserEligibility makes an HTTP call to UserService to check eligibility
func (s *BookApplicationService) checkUserEligibility(userId string) (bool, error) {
	resp, err := http.Get(fmt.Sprintf("%s/users/%s/eligibility", s.gatewayServiceURL, userId))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to check eligibility, status: %s", resp.Status)
	}

	var eligibilityResp map[string]bool
	if err := json.NewDecoder(resp.Body).Decode(&eligibilityResp); err != nil {
		return false, err
	}

	return eligibilityResp["isEligible"], nil
}
