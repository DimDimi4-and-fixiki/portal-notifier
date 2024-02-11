package usecase

import e "notify/internal/entity"

type authService interface {
	CreateUserWithDetails(info e.UserCommonInfo, cred e.UserCred) (e.User, error)
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
