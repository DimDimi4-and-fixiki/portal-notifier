package usecase

import (
	"context"
	"fmt"
	e "notify/internal/entity"
	"time"
)

func (u *UseCase) JWTFromUser(ctx context.Context, user *e.UserDB) (*e.JWTString, error) {
	return u.authService.JWTFromUser(ctx, user)
}

func (u *UseCase) VerifyJWT(ctx context.Context, token string) (bool, error) {
	user, err := u.authService.DecodeJWTToUser(ctx, token)
	if err != nil {
		return false, err
	}
	expTime, err := user.GetExpirationTime()
	if err != nil {
		return false, fmt.Errorf("error verifying JWT: %v", err)
	}
	if expTime.Before(time.Now()) {
		return false, fmt.Errorf("jwt token is expired, generate new token")
	}

	return true, nil
}

func (u *UseCase) JWTFromUserLogin(ctx context.Context, login string) (*e.JWTString, error) {
	return u.authService.JWTFromLogin(ctx, login)
}

func (u *UseCase) DecodeJWTToUser(ctx context.Context, token string) (*e.UserJWT, error) {
	return u.authService.DecodeJWTToUser(ctx, token)
}
