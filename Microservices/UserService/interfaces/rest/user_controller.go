package rest

import (
	"UserService/application/commands"
	"UserService/application/queries"
	"UserService/application/services"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type UserController struct {
	service *services.UserApplicationService
}

func NewUserController(service *services.UserApplicationService) *UserController {
	return &UserController{service: service}
}

// RegisterUser godoc
// @Summary Register a new user
// @Description Register a new user with the provided details
// @Tags users
// @Accept json
// @Produce json
// @Param user body commands.RegisterUserCommand true "Register user command"
// @Success 201 {object} entities.User
// @Failure 400 {object} map[string]string
// @Router /users [post]
func (c *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var cmd commands.RegisterUserCommand
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	user, err := c.service.RegisterUser(cmd)
	if err != nil {
		http.Error(w, "Error registering user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Retrieve a list of all users
// @Tags users
// @Produce json
// @Success 200 {array} entities.User
// @Failure 500 {object} map[string]string
// @Router /users [get]
func (c *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.service.GetAllUsers()
	if err != nil {
		http.Error(w, "Error retrieving users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

// GetUser godoc
// @Summary Get a user by ID
// @Description Get a user by their unique ID
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} entities.User
// @Failure 404 {object} map[string]string
// @Router /users/{id} [get]
func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	query := queries.GetUserQuery{UserID: id}
	user, err := c.service.GetUser(query)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
