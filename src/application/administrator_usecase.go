package application

import (
	"context"

	"github.com/gaku3601/ddd-golang/src/domain/model"
	"github.com/gaku3601/ddd-golang/src/domain/repository"
)

type AdministratorUseCase interface {
	RegisterAdministrator(ctx context.Context, dto *CreateAdministratorDto) error
}

type administratorUseCase struct {
	userRepository repository.AdministratorRepository
}

func NewAdministratorUseCase(repository repository.AdministratorRepository) AdministratorUseCase {
	return &administratorUseCase{repository}
}

// 管理者を登録する
func (u *administratorUseCase) RegisterAdministrator(ctx context.Context, dto *CreateAdministratorDto) error {
	userEmail, err := model.NewAdministratorEmail(dto.Email)
	if err != nil {
		return err
	}
	administrator := model.NewRegisterAdministrator(userEmail)
	return u.userRepository.Save(ctx, administrator)
}
