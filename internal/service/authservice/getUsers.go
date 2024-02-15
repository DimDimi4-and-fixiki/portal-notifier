package authservice

import (
	"context"
	"github.com/google/uuid"
	e "notify/internal/entity"
)

func (s *Service) GetAllUsers(ctx context.Context) (*[]e.UserDB, error) {
	return s.userRepo.GetAll(ctx)
}

// GetUser returns user from DB by login or by user_id
func (s *Service) GetUser(ctx context.Context, data e.GetUserInput) (*e.UserDB, error) {
	if data == (e.GetUserInput{}) {
		return nil, ErrEmptyUserInput
	}

	if data.ID != (uuid.UUID{}) {
		return s.userRepo.GetByID(ctx, data.ID)
	}
	return s.userRepo.GetByLogin(ctx, data.Login)

}

func (s *Service) GetUserByLogin(ctx context.Context, login string) (*e.UserDB, error) {
	return s.userRepo.GetByLogin(ctx, login)
}

func (s *Service) GetUserByID(ctx context.Context, id uuid.UUID) (*e.UserDB, error) {
	return s.userRepo.GetByID(ctx, id)
}
