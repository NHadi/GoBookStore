package repositories

import (
	"BookStore/domain/entities"
	"log"
	"time"

	"github.com/gocql/gocql"
)

type CassandraBookRepository struct {
	session *gocql.Session
}

// NewCassandraBookRepository creates a new instance of CassandraBookRepository
func NewCassandraBookRepository() (*CassandraBookRepository, error) {
	cluster := gocql.NewCluster("cassandra")
	cluster.Port = 9042
	cluster.Keyspace = "bookstore"

	var session *gocql.Session
	var err error
	retries := 10
	for i := 0; i < retries; i++ {
		session, err = cluster.CreateSession()
		if err == nil {
			log.Println("Connected to Cassandra!")
			return &CassandraBookRepository{session: session}, nil
		}
		log.Printf("Failed to connect to Cassandra, retrying in 5 seconds... (Attempt %d/%d)\n", i+1, retries)
		time.Sleep(5 * time.Second) // wait before retrying
	}
	return nil, err
}

// FindByID fetches a book by its ID from Cassandra
func (r *CassandraBookRepository) FindById(id string) (*entities.Book, error) {
	var book entities.Book
	err := r.session.Query("SELECT id, title, author FROM books WHERE id=?", id).Consistency(gocql.One).Scan(&book.ID, &book.Title, &book.Author.Name)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

// Add inserts a new book into Cassandra
func (r *CassandraBookRepository) Add(book *entities.Book) {
	err := r.session.Query("INSERT INTO books (id, title, author) VALUES (?, ?, ?)", book.ID, book.Title, book.Author.Name).Exec()
	if err != nil {
		log.Printf("Error inserting book into Cassandra: %v", err)
	}
}

// List retrieves all books from Cassandra
func (r *CassandraBookRepository) List() []*entities.Book {
	var books []*entities.Book
	iter := r.session.Query("SELECT id, title, author FROM books").Iter()
	defer iter.Close()

	var book entities.Book
	for iter.Scan(&book.ID, &book.Title, &book.Author.Name) {
		books = append(books, &book)
	}
	if err := iter.Close(); err != nil {
		// Handle the error appropriately
		log.Printf("Error iterating over books: %v", err)
	}
	return books
}

// Update modifies an existing book in Cassandra
func (r *CassandraBookRepository) Update(book *entities.Book) error {
	err := r.session.Query("UPDATE books SET title=?, author=? WHERE id=?", book.Title, book.Author.Name, book.ID).Exec()
	return err
}

// Delete removes a book from Cassandra
func (r *CassandraBookRepository) Delete(id string) error {
	err := r.session.Query("DELETE FROM books WHERE id=?", id).Exec()
	return err
}

// IndexBook adds a new book document to Cassandra
func (r *CassandraBookRepository) IndexBook(book *entities.Book) error {
	err := r.session.Query("INSERT INTO books (id, title, author) VALUES (?, ?, ?)", book.ID, book.Title, book.Author.Name).Exec()
	if err != nil {
		return err
	}
	return nil
}

// SearchBooks searches for books by title
func (r *CassandraBookRepository) SearchBooks(title string) ([]entities.Book, error) {
	var books []entities.Book

	// Using a prepared statement for better performance and security
	query := "SELECT id, title, author FROM books WHERE title LIKE ?"
	iter := r.session.Query(query, "%"+title+"%").Iter()
	defer iter.Close()

	var book entities.Book
	for iter.Scan(&book.ID, &book.Title, &book.Author) {
		books = append(books, book)
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return books, nil
}
