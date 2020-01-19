package model

type User struct {
	uid   *UserUID
	email *UserEmail
}

func NewUser(uid *UserUID, email *UserEmail) *User {
	return &User{uid: uid, email: email}
}

func (u *User) UserUID() *UserUID {
	return u.uid
}
func (u *User) UserName() *UserEmail {
	return u.email
}

type UserUID struct {
	value string
}

func NewUserUID(value string) *UserUID {
	return &UserUID{value: value}
}

func (u *UserUID) Value() string {
	return u.value
}

type UserEmail struct {
	value string
}

func NewUserEmail(value string) (*UserEmail, error) {
	return &UserEmail{value: value}, nil
}

func (u *UserEmail) Value() string {
	return u.value
}
