package services

import (
	"UserService/application/commands"
	"UserService/application/queries"
	"UserService/domain/entities"
	"UserService/domain/repositories"
	"errors"
)

type UserApplicationService struct {
	repo repositories.UserRepository
}

func NewUserApplicationService(repo repositories.UserRepository) *UserApplicationService {
	return &UserApplicationService{repo: repo}
}

func (s *UserApplicationService) CheckUserEligibility(userID string) (bool, error) {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return false, err
	}
	return user.IsEligible, nil
}

func (s *UserApplicationService) GetAllUsers() ([]*entities.User, error) {
	return s.repo.List()
}

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

func (s *UserApplicationService) GetUser(query queries.GetUserQuery) (*entities.User, error) {
	return s.repo.FindByID(query.UserID)
}
