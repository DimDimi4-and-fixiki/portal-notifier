package v1

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) Ping(ctx *gin.Context) {
	ctx.String(200, "Pong")
}
