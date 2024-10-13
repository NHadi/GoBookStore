// application/services/user_service_interface.go

package services

import (
	"UserService/application/commands"
	"UserService/application/queries"
	"UserService/domain/entities"
)

type UserApplicationServiceInterface interface {
	CheckUserEligibility(userID string) (bool, error)
	GetAllUsers() ([]*entities.User, error)
	RegisterUser(cmd commands.RegisterUserCommand) (*entities.User, error)
	GetUser(query queries.GetUserQuery) (*entities.User, error)
}
