package usecase

import (
	"context"
	e "notify/internal/entity"
)

func (u *UseCase) GetUser(ctx context.Context, data e.GetUserReq) (e.UserDB, error) {
	user, err := u.authService.GetUser(ctx, data)
	if err != nil {
		return e.UserDB{}, err
	}
	return *user, nil
}
