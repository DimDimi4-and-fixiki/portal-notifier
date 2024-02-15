package usecase

import (
	"context"
	e "notify/internal/entity"
)

func (u *UseCase) GetAllUsers(ctx context.Context) (*[]e.UserDB, error) {
	return u.authService.GetAllUsers(ctx)
}
