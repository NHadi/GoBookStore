// tests/mocks/mock_user_application_service.go

package mocks

import (
	"UserService/application/commands"
	"UserService/application/queries"
	"UserService/domain/entities"
)

type MockUserApplicationService struct {
	CheckUserEligibilityFunc func(userID string) (bool, error)
	RegisterUserFunc         func(cmd commands.RegisterUserCommand) (*entities.User, error)
	GetAllUsersFunc          func() ([]*entities.User, error)
	GetUserFunc              func(query queries.GetUserQuery) (*entities.User, error)
}

func (m *MockUserApplicationService) CheckUserEligibility(userID string) (bool, error) {
	if m.CheckUserEligibilityFunc != nil {
		return m.CheckUserEligibilityFunc(userID)
	}
	return false, nil // or a suitable default value
}

func (m *MockUserApplicationService) RegisterUser(cmd commands.RegisterUserCommand) (*entities.User, error) {
	if m.RegisterUserFunc != nil {
		return m.RegisterUserFunc(cmd)
	}
	return nil, nil // or a suitable default value
}

func (m *MockUserApplicationService) GetAllUsers() ([]*entities.User, error) {
	if m.GetAllUsersFunc != nil {
		return m.GetAllUsersFunc()
	}
	return nil, nil // or a suitable default value
}

func (m *MockUserApplicationService) GetUser(query queries.GetUserQuery) (*entities.User, error) {
	if m.GetUserFunc != nil {
		return m.GetUserFunc(query)
	}
	return nil, nil // or a suitable default value
}
