package swagger

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"notify/pkg/logger"
)

type Handler struct {
	logger logger.Logger
	router *gin.Engine
}

// NewSwagger creates new handler for swagger docs
func NewSwagger(logs logger.Logger, router *gin.Engine) *Handler {
	return &Handler{logger: logs, router: router}
}

// InitDocs binds route for displaying swagger docs to handler
func (h *Handler) InitDocs() {
	docsGroup := h.router.Group("/docs")
	{
		docsGroup.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
