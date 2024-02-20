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
	GetUser(ctx context.Context, data e.GetUserReq) (*e.UserDB, error)
}

type notifyService interface {
	SendProposalToOwner(ctx context.Context, proposal *e.ProjectProposal, ownerEmail string) error
	SendProposalToUser(ctx context.Context, proposal *e.ProjectProposal, userEmail string) error
	SaveProposal(ctx context.Context, proposal *e.ProjectProposal) error
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
