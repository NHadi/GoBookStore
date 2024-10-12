package event_handlers

import (
	"BookStore/domain/entities"
	"BookStore/domain/events"
)

// HandleBookAddedEvent handles the event of a book being added
func HandleBookAddedEvent(book entities.Book) {
	event := events.BookAddedEvent{Book: book}
	event.Notify() // Here, you could trigger external processes, like sending notifications
}
