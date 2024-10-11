package domain

// Book represents a book entity
type Book struct {
	ID     string
	Title  string
	Author Author // Use Author as a Value Object
}
