package application_test

import (
	"UserService/application/commands"
	"UserService/application/services"
	"UserService/domain/entities"
	"UserService/tests/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserApplicationService_RegisterUser_Success(t *testing.T) {
	mockRepo := &mocks.MockUserRepository{Users: make(map[string]*entities.User)}
	userService := services.NewUserApplicationService(mockRepo)

	cmd := commands.RegisterUserCommand{Name: "Alice", Email: "alice@example.com", Password: "password", IsEligible: true}
	user, err := userService.RegisterUser(cmd)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "Alice", user.Name)
	assert.Equal(t, "alice@example.com", user.Email)
}

func TestUserApplicationService_RegisterUser_UserAlreadyExists(t *testing.T) {
	mockRepo := &mocks.MockUserRepository{Users: make(map[string]*entities.User)}
	userService := services.NewUserApplicationService(mockRepo)

	existingUser := &entities.User{Email: "existing@example.com"}
	mockRepo.Users["existing@example.com"] = existingUser

	cmd := commands.RegisterUserCommand{Name: "Bob", Email: "existing@example.com", Password: "password", IsEligible: true}
	user, err := userService.RegisterUser(cmd)
	assert.Error(t, err)
	assert.Nil(t, user)
}
