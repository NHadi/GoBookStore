package domain

// BookRepository defines methods for book operations
type BookRepository interface {
	Add(book Book)
	List() []Book
	Update(book Book)
	Delete(id string)
	FindByID(id string) (Book, bool)
}
