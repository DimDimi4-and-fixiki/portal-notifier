package admin

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	e "notify/internal/entity"
	"notify/pkg/logger"
)

type UseCase interface {
	GetAllUsers(ctx context.Context) (*[]e.UserDB, error)
	UpdateUserInfo(ctx context.Context, id uuid.UUID, data e.UserCommonInfo) (e.UserDB, error)
	CreateUser(ctx context.Context, user *e.User) (e.UserDB, error)
	GetUser(ctx context.Context, data e.GetUserReq) (e.UserDB, error)
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
	admin := h.router.Group("/api/admin")
	{
		admin.GET("/ping", h.Ping)
		admin.POST("/get_user", h.getUser)
		admin.POST("/register_user", h.RegisterUser)
	}
}
