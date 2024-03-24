package usecase

import (
	"context"
	"fmt"
	e "notify/internal/entity"
	"notify/pkg/crypt"
)

func (u *UseCase) AuthUserApiToken(ctx context.Context, login string, password string) (bool, error) {
	data := e.GetUserInput{Login: login}
	user, err := u.authService.GetUser(ctx, data)
	if err != nil {
		return false, err
	}
	if user.Info.Role != e.Service {
		return false, fmt.Errorf("api_token auth method is only for services, your role %s", user.Info.Role)
	}

	return crypt.VerifyPassword(password, user.Cred.HashedPassword), nil
}
