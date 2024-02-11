package usecase

import e "notify/internal/entity"

func (u *UseCase) CreateUserWithDetails(info e.UserCommonInfo, cred e.UserCred) (e.User, error) {
	return u.authService.CreateUserWithDetails(info, cred)
}
