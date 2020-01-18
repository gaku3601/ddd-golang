package infrastructure

import (
	"context"
	"fmt"

	"github.com/gaku3601/ddd-golang/src/domain/model"
)

type UserRepository struct{}

func (u *UserRepository) Save(ctx context.Context, user *model.User) error {
	fmt.Println(ctx)
	fmt.Println(user.UserUID().Value())
	fmt.Println(user.UserName().Value())
	return nil
}
