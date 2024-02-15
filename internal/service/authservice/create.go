package authservice

import (
	"context"
	e "notify/internal/entity"
)

func (s *Service) CreateUser(ctx context.Context, user *e.User) (e.UserDB, error) {
	res, err := s.userRepo.Create(ctx, user)
	if err != nil {
		return e.UserDB{}, err
	}
	return *res, nil
}
