package repositories

import (
	"UserService/domain/entities"
	"errors"

	"github.com/google/uuid"
)

type InMemoryUserRepository struct {
	users map[string]*entities.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{users: make(map[string]*entities.User)}
}

func (r *InMemoryUserRepository) FindByID(id string) (*entities.User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (r *InMemoryUserRepository) List() ([]*entities.User, error) {
	users := make([]*entities.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}
	return users, nil
}

func (r *InMemoryUserRepository) FindByEmail(email string) (*entities.User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, nil
}

func (r *InMemoryUserRepository) CreateUser(user *entities.User) error {
	// Generate a unique ID for the user
	user.ID = generateUniqueID()
	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) UpdateUser(user *entities.User) error {
	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) DeleteUser(id string) error {
	delete(r.users, id)
	return nil
}

func generateUniqueID() string {
	return uuid.New().String() // Generate a new UUID
}
