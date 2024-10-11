package domain

type BookService struct {
	repo BookRepository
}

// NewBookService creates a new BookService
func NewBookService(repo BookRepository) *BookService {
	return &BookService{repo: repo}
}

// AddBook adds a new book to the repository
func (s *BookService) AddBook(book Book) {
	s.repo.Add(book)
	event := BookAddedEvent{Book: book}
	event.Notify() // Notify the event
}

// ListBooks returns all books from the repository
func (s *BookService) ListBooks() []Book {
	return s.repo.List()
}

// UpdateBook updates an existing book in the repository
func (s *BookService) UpdateBook(book Book) {
	s.repo.Update(book)
}

// DeleteBook removes a book from the repository by ID
func (s *BookService) DeleteBook(id string) {
	s.repo.Delete(id)
}
