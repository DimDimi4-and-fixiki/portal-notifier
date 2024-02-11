package admin

import (
	"github.com/gin-gonic/gin"
	e "notify/internal/entity"
	"notify/pkg/logger"
)

type UseCase interface {
	CreateUserWithDetails(info e.UserCommonInfo, cred e.UserCred) (e.User, error)
}

type Handler struct {
	uc     UseCase
	logger logger.Logger
	router *gin.Engine
}

func NewHandler(uc UseCase, logs logger.Logger, router *gin.Engine) *Handler {
	return &Handler{uc: uc, logger: logs, router: router}
}

// WithRoutes Adds admin namespace routes to gin HTTP Router
func (h *Handler) WithRoutes() {
	admin := h.router.Group("/admin")
	{
		admin.GET("/ping", h.Ping)
	}
}
