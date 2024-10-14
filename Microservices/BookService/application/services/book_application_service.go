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
	cassandraRepo     repositories.BookRepository
	elasticsearchRepo repositories.BookRepository
	gatewayServiceURL string
}

func NewBookApplicationService(cassandraRepo, elasticsearchRepo repositories.BookRepository, gatewayServiceURL string) *BookApplicationService {
	return &BookApplicationService{
		cassandraRepo:     cassandraRepo,
		elasticsearchRepo: elasticsearchRepo,
		gatewayServiceURL: gatewayServiceURL,
	}
}

func (s *BookApplicationService) AddBook(cmd commands.AddBookCommand) (*entities.Book, error) {
	book, err := cmd.Handle(s.cassandraRepo)
	if err != nil {
		return nil, err
	}

	if err := s.elasticsearchRepo.IndexBook(book); err != nil {
		return nil, err
	}

	return book, nil
}

func (s *BookApplicationService) ListBooks(q queries.GetBooksQuery) []*entities.Book {
	return q.Handle(s.elasticsearchRepo)
}

func (s *BookApplicationService) FindByID(id string) (*entities.Book, error) {
	book, err := s.cassandraRepo.FindById(id)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (s *BookApplicationService) Update(book *entities.Book) error {
	if err := s.cassandraRepo.Update(book); err != nil {
		return err
	}

	if err := s.elasticsearchRepo.IndexBook(book); err != nil {
		return err
	}

	return nil
}

func (s *BookApplicationService) Delete(id string) error {
	if err := s.cassandraRepo.Delete(id); err != nil {
		return err
	}

	return nil
}

func (s *BookApplicationService) SearchBooks(title string) ([]entities.Book, error) {
	return s.elasticsearchRepo.SearchBooks(title)
}

func (s *BookApplicationService) BorrowBook(userId string, bookId string) (string, error) {
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
