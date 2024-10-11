package domain

import "fmt"

// BookAddedEvent represents an event when a book is added
type BookAddedEvent struct {
	Book Book
}

// Notify handles the event when a book is added
func (e BookAddedEvent) Notify() {
	fmt.Printf("Book Added: ID: %s, Title: %s, Author: %s\n", e.Book.ID, e.Book.Title, e.Book.Author.Name)
}
