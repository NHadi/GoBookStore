package repository

import "BookStore/domain" // Update with your actual module path

// InMemoryBookRepository is an in-memory implementation of BookRepository
type InMemoryBookRepository struct {
	books map[string]domain.Book
}

// NewInMemoryBookRepository creates a new InMemoryBookRepository
func NewInMemoryBookRepository() *InMemoryBookRepository {
	return &InMemoryBookRepository{
		books: make(map[string]domain.Book),
	}
}

// Add adds a new book to the repository
func (r *InMemoryBookRepository) Add(book domain.Book) {
	r.books[book.ID] = book
}

// List returns all books in the repository
func (r *InMemoryBookRepository) List() []domain.Book {
	var bookList []domain.Book
	for _, book := range r.books {
		bookList = append(bookList, book)
	}
	return bookList
}

// Update updates an existing book in the repository
func (r *InMemoryBookRepository) Update(book domain.Book) {
	r.books[book.ID] = book
}

// Delete removes a book from the repository by ID
func (r *InMemoryBookRepository) Delete(id string) {
	delete(r.books, id)
}

// FindByID retrieves a book by ID
func (r *InMemoryBookRepository) FindByID(id string) (domain.Book, bool) {
	book, exists := r.books[id]
	return book, exists
}
