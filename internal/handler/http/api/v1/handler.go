package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"notify/internal/config"
	e "notify/internal/entity"
	"notify/pkg/logger"
)

type UseCase interface {
	SendProposal(ctx context.Context, proposal *e.ProjectProposal, ownerEmail string) (*e.ProjectProposal, []error)
}

type Handler struct {
	uc     UseCase
	logger logger.Logger
	router *gin.Engine
	cfg    *config.Config
}

func NewHandler(uc UseCase, logs logger.Logger, router *gin.Engine, cfg *config.Config) *Handler {
	return &Handler{uc: uc, logger: logs, router: router, cfg: cfg}
}

// WithRoutes Adds v1 namespace routes to gin HTTP Router
func (h *Handler) WithRoutes() {
	v1 := h.router.Group("/api/v1")
	{
		v1.GET("/ping", h.ping)
		v1.POST("/send_proposal", h.sendProposal)
	}
}
