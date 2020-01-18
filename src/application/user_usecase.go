package application

import (
	"context"

	"github.com/gaku3601/ddd-golang/src/domain/model"
	"github.com/gaku3601/ddd-golang/src/domain/repository"
)

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(repository repository.UserRepository) *userUseCase {
	return &userUseCase{userRepository: repository}
}

func (u *userUseCase) RegisterUser(ctx context.Context, uid string, name string) error {
	userName, err := model.NewUserName(name)
	if err != nil {
		return err
	}
	userUID := model.NewUserUID(uid)
	user := model.NewUser(userUID, userName)
	return u.userRepository.Save(ctx, user)
}
