package repository

import (
	"context"

	"github.com/gaku3601/ddd-golang/src/domain/model"
)

type AdministratorRepository interface {
	Save(ctx context.Context, administrator *model.Administrator) error
}
