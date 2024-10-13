package mocks

import (
	"UserService/domain/entities"
	"errors"
)

// MockUserRepository is a mock implementation of the UserRepository interface
type MockUserRepository struct {
	Users map[string]*entities.User // make this field public to access it in tests
	Err   error                     // to simulate errors
}

// NewMockUserRepository initializes the mock repository
func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{Users: make(map[string]*entities.User)}
}

// FindByID returns a user by ID or an error if not found
func (m *MockUserRepository) FindByID(id string) (*entities.User, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	user, exists := m.Users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// FindByEmail returns a user by email or nil if not found
func (m *MockUserRepository) FindByEmail(email string) (*entities.User, error) {
	for _, user := range m.Users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, nil
}

// CreateUser adds a new user to the mock repository
func (m *MockUserRepository) CreateUser(user *entities.User) error {
	if m.Err != nil {
		return m.Err
	}
	m.Users[user.ID] = user
	return nil
}

// UpdateUser updates an existing user in the mock repository
func (m *MockUserRepository) UpdateUser(user *entities.User) error {
	if m.Err != nil {
		return m.Err
	}
	m.Users[user.ID] = user
	return nil
}

// DeleteUser removes a user from the mock repository
func (m *MockUserRepository) DeleteUser(id string) error {
	if m.Err != nil {
		return m.Err
	}
	delete(m.Users, id)
	return nil
}

// List returns all users in the mock repository
func (m *MockUserRepository) List() ([]*entities.User, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	users := make([]*entities.User, 0, len(m.Users))
	for _, user := range m.Users {
		users = append(users, user)
	}
	return users, nil
}
