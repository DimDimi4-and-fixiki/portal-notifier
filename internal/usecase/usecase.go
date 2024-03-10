package usecase

import (
	"context"
	"github.com/google/uuid"
	e "notify/internal/entity"
)

type authService interface {
	CreateUser(ctx context.Context, user *e.User) (e.UserDB, error)
	UpdateUserInfo(ctx context.Context, id uuid.UUID, data e.UserCommonInfo) (e.UserDB, error)
	Verify(ctx context.Context, u *e.User) (bool, error)
	GetAllUsers(ctx context.Context) (*[]e.UserDB, error)
	GetUser(ctx context.Context, data e.GetUserInput) (*e.UserDB, error)
	JWTFromUser(context.Context, *e.UserDB) (*e.JWTString, error)
	DecodeJWTToUser(context.Context, string) (*e.UserJWT, error)
	JWTFromLogin(ctx context.Context, login string) (*e.JWTString, error)
}

type notifyService interface {
	// TODO: define interface to inject a service
}

type UseCase struct {
	authService   authService
	notifyService notifyService
}

func NewUseCase(authService authService, notifyService notifyService) *UseCase {
	return &UseCase{
		authService:   authService,
		notifyService: notifyService,
	}
}
