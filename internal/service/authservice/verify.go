package authservice

import (
	"context"
	"errors"
	"fmt"
	e "notify/internal/entity"
	"notify/pkg/crypt"
)

var ErrVerifyUser = errors.New("verify user error")

func (s *Service) Verify(ctx context.Context, u *e.User) (bool, error) {
	user, err := s.userRepo.GetByLogin(ctx, u.Cred.Login)
	if err != nil {
		return false, fmt.Errorf("%w %w", err, ErrVerifyUser)
	}

	return crypt.VerifyPassword(u.Cred.Password, user.Cred.HashedPassword), nil
}
