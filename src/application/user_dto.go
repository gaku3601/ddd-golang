package application

type CreateUserDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
