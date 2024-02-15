package admin

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	e "notify/internal/entity"
)

func (h *Handler) getUser(c *gin.Context) {
	var input e.GetUserInput
	err := c.BindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.logger.Info("Parsed input")
	user, err := h.uc.GetUser(context.Background(), input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": user})
}
