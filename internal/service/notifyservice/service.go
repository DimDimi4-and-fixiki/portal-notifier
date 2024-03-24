package notifyservice

import (
	"context"
	courierc "github.com/trycourier/courier-go/v3"
	"go.uber.org/zap"
	e "notify/internal/entity"
	"notify/pkg/logger"
)

type userRepo interface {
	// TODO: define interface to inject a service or an adapter
}

type proposalRepo interface {
	Create(ctx context.Context, proposal *e.ProjectProposalDB) (*e.ProjectProposalDB, error)
}

type courierClient interface {
	SendTextMessage(ctx context.Context, msg e.EmailMsg) (*courierc.SendMessageResponse, error)
}

type Service struct {
	userRepo      userRepo
	proposalRepo  proposalRepo
	courierClient courierClient
	logger        *zap.SugaredLogger
}

func NewService(userRepo userRepo, proposalRepo proposalRepo, courierClient courierClient) *Service {
	return &Service{
		userRepo:      userRepo,
		proposalRepo:  proposalRepo,
		courierClient: courierClient,
		logger:        logger.NewSugar(),
	}
}
