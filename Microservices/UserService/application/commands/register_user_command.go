package commands

type RegisterUserCommand struct {
	Name       string
	Email      string
	Password   string
	IsEligible bool
}
