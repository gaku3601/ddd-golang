package infrastructure

import (
	"context"
	"fmt"

	"github.com/gaku3601/ddd-golang/src/domain/model"
)

type UserRepository struct{}

func (u *UserRepository) Save(ctx context.Context, userEmail *model.UserEmail, password string) error {
	fmt.Println(userEmail.Value())
	fmt.Println(password)
	return nil
}
