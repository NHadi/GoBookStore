package main

import (
	"BookStore/domain"     // Update with your actual module path
	"BookStore/repository" // Update with your actual module path
	"fmt"
)

func main() {
	bookRepo := repository.NewInMemoryBookRepository()
	bookService := domain.NewBookService(bookRepo)

	// Create Authors
	author1, _ := domain.NewAuthor("Nurul Hadi")
	author2, _ := domain.NewAuthor("Jane Doe")

	// Create Books
	book1 := domain.Book{ID: "1", Title: "Go Programming", Author: author1}
	bookService.AddBook(book1)

	book2 := domain.Book{ID: "2", Title: "Learn Go", Author: author2}
	bookService.AddBook(book2)

	// Read
	books := bookService.ListBooks()
	fmt.Println("Books in the repository:")
	for _, book := range books {
		fmt.Printf("ID: %s, Title: %s, Author: %s\n", book.ID, book.Title, book.Author.Name)
	}

	// Update
	book2.Title = "Learning Go"
	bookService.UpdateBook(book2)

	// Read updated books
	updatedBooks := bookService.ListBooks()
	fmt.Println("\nBooks after update:")
	for _, book := range updatedBooks {
		fmt.Printf("ID: %s, Title: %s, Author: %s\n", book.ID, book.Title, book.Author.Name)
	}

	// Delete
	bookService.DeleteBook("1")

	// Read books after deletion
	booksAfterDeletion := bookService.ListBooks()
	fmt.Println("\nBooks after deletion:")
	for _, book := range booksAfterDeletion {
		fmt.Printf("ID: %s, Title: %s, Author: %s\n", book.ID, book.Title, book.Author.Name)
	}
}
