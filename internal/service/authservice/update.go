package authservice

import (
	"context"
	"github.com/google/uuid"
	e "notify/internal/entity"
)

func (s *Service) UpdateUserInfo(ctx context.Context, id uuid.UUID, data e.UserCommonInfo) (e.UserDB, error) {
	res, err := s.userRepo.UpdateCommonInfoByID(ctx, id, data)
	return *res, err
}
