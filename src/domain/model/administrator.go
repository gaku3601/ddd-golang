package model

type Administrator struct {
	email *AdministratorEmail
}

func NewRegisterAdministrator(email *AdministratorEmail) *Administrator {
	return &Administrator{email: email}
}

func (u *Administrator) AdministratorEmail() *AdministratorEmail {
	return u.email
}

type AdministratorEmail struct {
	value string
}

func NewAdministratorEmail(value string) (*AdministratorEmail, error) {
	return &AdministratorEmail{value: value}, nil
}

func (u *AdministratorEmail) Value() string {
	return u.value
}
