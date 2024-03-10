package usecase

import (
	"context"
	e "notify/internal/entity"
)

func userToServiceUser(u *e.User) e.ServiceUser {
	return e.ServiceUser{
		Cred: e.ServiceCred{
			Login:  u.Cred.Login,
			ApiKey: u.Cred.Password,
		},
		Info: e.ServiceUserCommonInfo{
			Name: u.Info.Name,
		},
	}
}

func (u *UseCase) CreateService(ctx context.Context, user *e.User) (e.ServiceUser, error) {
	_, err := u.authService.CreateUser(ctx, user)
	if err != nil {
		return e.ServiceUser{}, err
	}
	return userToServiceUser(user), nil
}
