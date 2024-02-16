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

func NewHandler(logs logger.Logger, router *gin.Engine) *Handler {
	return &Handler{logger: logs, router: router}
}

// InitDocs binds route for swagger docs
func (h *Handler) InitDocs() {
	docsGroup := h.router.Group("/docs")
	{
		docsGroup.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
