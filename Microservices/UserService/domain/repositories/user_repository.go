package repositories

import "UserService/domain/entities"

type UserRepository interface {
	FindByID(id string) (*entities.User, error)
	FindByEmail(email string) (*entities.User, error)
	CreateUser(user *entities.User) error
	UpdateUser(user *entities.User) error
	DeleteUser(id string) error
	List() ([]*entities.User, error)
}
