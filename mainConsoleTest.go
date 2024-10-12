package main

import (
	"BookStore/application/commands"
	"BookStore/application/queries"
	"BookStore/application/services"
	"BookStore/infrastructure/event_handlers"
	"BookStore/infrastructure/repositories"
	"fmt"
)

func mainConsoleTest() {
	// Initialize repository and application service
	bookRepo := repositories.NewInMemoryBookRepository()
	bookService := services.NewBookApplicationService(bookRepo)

	// Add books using a command
	addBookCommand1 := commands.AddBookCommand{ID: "1", Title: "Go Programming", Author: "Nurul Hadi"}
	book1, _ := bookService.AddBook(addBookCommand1)

	addBookCommand2 := commands.AddBookCommand{ID: "2", Title: "Advanced Go", Author: "Jane Doe"}
	book2, _ := bookService.AddBook(addBookCommand2)

	// List books using a query
	query := queries.GetBooksQuery{}
	books := bookService.ListBooks(query)

	fmt.Println("Books in the system:")
	for _, book := range books {
		fmt.Printf("ID: %s, Title: %s, Author: %s\n", book.ID, book.Title, book.Author.Name)
	}

	// Handle domain events
	event_handlers.HandleBookAddedEvent(*book1)
	event_handlers.HandleBookAddedEvent(*book2)
}
