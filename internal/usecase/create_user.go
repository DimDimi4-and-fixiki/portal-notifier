package usecase

import (
	"context"
	e "notify/internal/entity"
)

func (u *UseCase) CreateUser(ctx context.Context, user *e.User) (e.UserDB, error) {
	res, err := u.authService.CreateUser(ctx, user)
	if err != nil {
		return e.UserDB{}, err
	}
	return res, nil
}
