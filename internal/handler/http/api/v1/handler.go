package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	e "notify/internal/entity"
	"notify/pkg/logger"
)

type UseCase interface {
	SendProposal(ctx context.Context) (interface{}, error)
	Pong(ctx context.Context) (interface{}, error)
	CreateUser(ctx context.Context, user *e.User) (e.UserDB, error)
}

type Handler struct {
	uc     UseCase
	logger logger.Logger
	router *gin.Engine
}

func NewHandler(uc UseCase, logs logger.Logger, router *gin.Engine) *Handler {
	return &Handler{uc: uc, logger: logs, router: router}
}

// WithRoutes Adds v1 namespace routes to gin HTTP Router
func (h *Handler) WithRoutes() {
	v1 := h.router.Group("/v1")
	{
		v1.GET("/ping", h.Ping)
	}
}
