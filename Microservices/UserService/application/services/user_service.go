package services

import (
	"UserService/application/commands"
	"UserService/application/queries"
	"UserService/domain/entities"
	"UserService/domain/repositories"
	"errors"
)

// UserApplicationService implements UserApplicationServiceInterface
type UserApplicationService struct {
	repo repositories.UserRepository
}

// NewUserApplicationService creates a new instance of UserApplicationService
func NewUserApplicationService(repo repositories.UserRepository) *UserApplicationService {
	return &UserApplicationService{repo: repo}
}

// CheckUserEligibility checks if a user is eligible
func (s *UserApplicationService) CheckUserEligibility(userID string) (bool, error) {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return false, err
	}
	return user.IsEligible, nil
}

// GetAllUsers retrieves all users
func (s *UserApplicationService) GetAllUsers() ([]*entities.User, error) {
	return s.repo.List()
}

// RegisterUser registers a new user
func (s *UserApplicationService) RegisterUser(cmd commands.RegisterUserCommand) (*entities.User, error) {
	existingUser, _ := s.repo.FindByEmail(cmd.Email)
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	newUser := &entities.User{
		Name:       cmd.Name,
		Email:      cmd.Email,
		Password:   cmd.Password, // Password should be hashed
		IsEligible: cmd.IsEligible,
	}

	if err := s.repo.CreateUser(newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}

// GetUser retrieves a user by ID
func (s *UserApplicationService) GetUser(query queries.GetUserQuery) (*entities.User, error) {
	return s.repo.FindByID(query.UserID)
}
