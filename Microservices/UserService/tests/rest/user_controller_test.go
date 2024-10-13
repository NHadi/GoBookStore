// tests/rest/user_controller_test.go

package rest_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"UserService/application/commands"
	"UserService/domain/entities"
	"UserService/interfaces/rest"
	"UserService/tests/mocks"

	"github.com/stretchr/testify/assert"
)

func TestUserController_RegisterUser(t *testing.T) {
	// Arrange
	mockService := &mocks.MockUserApplicationService{}
	mockService.RegisterUserFunc = func(cmd commands.RegisterUserCommand) (*entities.User, error) {
		return &entities.User{ID: "1", Name: cmd.Name, Email: cmd.Email}, nil
	}

	controller := rest.NewUserController(mockService)

	userCmd := commands.RegisterUserCommand{Name: "Alice", Email: "alice@example.com", Password: "securepassword", IsEligible: true}
	userJSON, _ := json.Marshal(userCmd)

	req, err := http.NewRequest("POST", "/users", bytes.NewReader(userJSON)) // Use bytes.NewReader
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	// Act
	http.HandlerFunc(controller.RegisterUser).ServeHTTP(rr, req)

	// Assert
	assert.Equal(t, http.StatusCreated, rr.Code)
	var userResponse entities.User
	err = json.Unmarshal(rr.Body.Bytes(), &userResponse)
	assert.NoError(t, err)
	assert.Equal(t, "Alice", userResponse.Name)
	assert.Equal(t, "alice@example.com", userResponse.Email)
}
