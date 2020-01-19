package application

import (
	"context"

	"github.com/gaku3601/ddd-golang/src/domain/model"
	"github.com/gaku3601/ddd-golang/src/domain/repository"
)

type UserUseCase interface {
	RegisterUser(ctx context.Context, uid string, name string) error
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(repository repository.UserRepository) UserUseCase {
	return &userUseCase{repository}
}

func (u *userUseCase) RegisterUser(ctx context.Context, uid string, email string) error {
	userEmail, err := model.NewUserEmail(email)
	if err != nil {
		return err
	}
	userUID := model.NewUserUID(uid)
	user := model.NewUser(userUID, userEmail)
	return u.userRepository.Save(ctx, user)
}
