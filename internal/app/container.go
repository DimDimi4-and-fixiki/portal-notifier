package app

import (
	"go.uber.org/zap"
	"gopkg.in/tylerb/graceful.v1"
	"gorm.io/gorm"
	"notify/internal/repository/pgrepo/proposalrepo"
	"notify/internal/repository/pgrepo/userrepo"
	"notify/internal/service/authservice"
	"notify/internal/service/notifyservice"
	"notify/internal/service/pingpong"
	"notify/internal/usecase"
)

type Container struct {
	pg         *gorm.DB
	httpServer *graceful.Server

	deps map[string]interface{}
}

func NewContainer(pg *gorm.DB, httpServer *graceful.Server) *Container {

	return &Container{
		pg:         pg,
		httpServer: httpServer,

		deps: make(map[string]interface{}),
	}
}

func (c *Container) GetUseCase() *usecase.UseCase {

	return usecase.NewUseCase(c.getAuthService(), c.getNotifyService())
}

func (c *Container) getPg() *gorm.DB {
	return c.pg
}

func (c *Container) getHttpServer() *graceful.Server {
	return c.httpServer
}

func (c *Container) getAuthService() *authservice.Service {

	return authservice.NewService(c.getUserRepo())
}

func (c *Container) getNotifyService() *notifyservice.Service {

	return notifyservice.NewService(c.getUserRepo(), c.getProposalRepo())
}

func (c *Container) getPingPong() *pingpong.Service {

	return pingpong.NewService()
}

func (c *Container) getUserRepo() *userrepo.Repository {

	repo, err := userrepo.New(c.getPg()).WithMigrate()
	if err != nil {
		a.logger.Fatal("Failed to migrate UserRepo Entities", zap.Error(err))
	}
	return repo
}

func (c *Container) getProposalRepo() *proposalrepo.Repository {

	repo, err := proposalrepo.New(c.getPg()).WithMigrate()
	if err != nil {
		a.logger.Fatal("Failed to migrate ProposalRepo Entities", zap.Error(err))
	}
	return repo
}
