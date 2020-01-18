package model

import "errors"

type User struct {
	uid  *UserUID
	name *UserName
}

func NewUser(uid *UserUID, name *UserName) *User {
	return &User{uid: uid, name: name}
}

func (u *User) UserUID() *UserUID {
	return u.uid
}
func (u *User) UserName() *UserName {
	return u.name
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

type UserName struct {
	value string
}

func NewUserName(value string) (*UserName, error) {
	if len(value) <= 2 {
		return nil, errors.New("名前は2文字以上が必要です")
	}
	return &UserName{value: value}, nil
}

func (u *UserName) Value() string {
	return u.value
}
