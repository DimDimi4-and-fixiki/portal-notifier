package admin

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	e "notify/internal/entity"
)

func (h *Handler) RegisterUser(c *gin.Context) {
	var user e.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		h.logger.Info("invalid request:", zap.Error(err))
		return
	}

	result, err := h.uc.CreateUser(context.Background(), &user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}
