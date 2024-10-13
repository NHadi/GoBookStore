package infrastructure_test

import (
	"UserService/domain/entities"
	"UserService/infrastructure/repositories"
	"testing"
)

func TestInMemoryUserRepository_CreateUser(t *testing.T) {
	repo := repositories.NewInMemoryUserRepository()
	user := &entities.User{Name: "John", Email: "john@example.com"}

	err := repo.CreateUser(user)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	fetchedUser, err := repo.FindByID(user.ID)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if fetchedUser.Name != user.Name {
		t.Errorf("expected %s, got %s", user.Name, fetchedUser.Name)
	}
}
