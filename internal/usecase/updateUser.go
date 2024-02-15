package usecase

import (
	"context"
	"github.com/google/uuid"
	e "notify/internal/entity"
)

func (u *UseCase) UpdateUserInfo(ctx context.Context, id uuid.UUID, data e.UserCommonInfo) (e.UserDB, error) {
	return u.authService.UpdateUserInfo(ctx, id, data)
}
