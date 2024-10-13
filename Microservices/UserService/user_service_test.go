package tests

import (
	"UserService/application/commands"
	"UserService/application/services"
	"UserService/domain/entities"
	"UserService/tests/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserApplicationService_CheckUserEligibility_UserIsEligible_ReturnsTrue(t *testing.T) {
	mockRepo := &mocks.MockUserRepository{Users: make(map[string]*entities.User)}
	userService := services.NewUserApplicationService(mockRepo)

	// Test with an eligible user
	mockRepo.Users["456"] = &entities.User{IsEligible: true}
	isEligible, err := userService.CheckUserEligibility("456")
	assert.NoError(t, err)
	assert.True(t, isEligible)
}

func TestUserApplicationService_CheckUserEligibility_UserNotFound_ReturnsError(t *testing.T) {
	mockRepo := &mocks.MockUserRepository{Users: make(map[string]*entities.User)}
	userService := services.NewUserApplicationService(mockRepo)

	// Test with a user not found
	mockRepo.Err = errors.New("user not found")
	isEligible, err := userService.CheckUserEligibility("999")
	assert.Error(t, err)
	assert.False(t, isEligible)
}

func TestUserApplicationService_RegisterUser_UserAlreadyExists_ReturnsError(t *testing.T) {
	mockRepo := &mocks.MockUserRepository{Users: make(map[string]*entities.User)}
	userService := services.NewUserApplicationService(mockRepo)

	mockRepo.Users["existing@example.com"] = &entities.User{Email: "existing@example.com"}
	cmd := commands.RegisterUserCommand{Name: "John", Email: "existing@example.com", Password: "password", IsEligible: true}

	user, err := userService.RegisterUser(cmd)
	assert.Error(t, err)
	assert.Nil(t, user)
}

func TestUserApplicationService_RegisterUser_Success(t *testing.T) {
	mockRepo := &mocks.MockUserRepository{Users: make(map[string]*entities.User)}
	userService := services.NewUserApplicationService(mockRepo)

	cmd := commands.RegisterUserCommand{Name: "John", Email: "john@example.com", Password: "password", IsEligible: true}
	user, err := userService.RegisterUser(cmd)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "John", user.Name)
	assert.Equal(t, "john@example.com", user.Email)
}
