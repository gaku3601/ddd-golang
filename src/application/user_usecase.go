package application

import (
	"context"

	"github.com/gaku3601/ddd-golang/src/domain/model"
	"github.com/gaku3601/ddd-golang/src/domain/repository"
)

type UserUseCase interface {
	RegisterUser(ctx context.Context, dto *CreateUserDto) error
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(repository repository.UserRepository) UserUseCase {
	return &userUseCase{repository}
}

func (u *userUseCase) RegisterUser(ctx context.Context, dto *CreateUserDto) error {
	userEmail, err := model.NewUserEmail(dto.Email)
	if err != nil {
		return err
	}
	return u.userRepository.Save(ctx, userEmail, dto.Password)
}
