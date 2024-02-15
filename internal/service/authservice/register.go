package authservice

import (
	"context"
	e "notify/internal/entity"
)

func (s *Service) RegisterUser(ctx context.Context, user *e.User) (e.UserDB, error) {
	res, err := s.userRepo.Create(ctx, user)
	return *res, err
}
