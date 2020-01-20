package repository

import (
	"context"

	"github.com/gaku3601/ddd-golang/src/domain/model"
)

type UserRepository interface {
	Save(ctx context.Context, email *model.UserEmail, password string) error
}
