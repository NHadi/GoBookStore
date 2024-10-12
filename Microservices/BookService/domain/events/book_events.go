package events

import "BookStore/domain/entities"

// BookAddedEvent represents an event triggered when a book is added
type BookAddedEvent struct {
	Book entities.Book
}

// Notify triggers the event
func (e BookAddedEvent) Notify() {
	// Trigger event (e.g., send to a message broker or log)
	println("Book Added:", e.Book.Title)
}
