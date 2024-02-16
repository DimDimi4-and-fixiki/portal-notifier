package admin

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	e "notify/internal/entity"
)

// RegisterUser godoc
// @Summary Register new user
// @Description Register new user with user details
// @Tags Admin
// @Accept application/json
// @Produce application/json
// @Param user body e.User true "User details"
// @Success 200 {object} e.HttpResp[e.UserCommonInfo]
// @Router /admin/register_user/ [post]
func (h *Handler) RegisterUser(c *gin.Context) {
	var request e.User
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		h.logger.Info("invalid request:", zap.Error(err))
		return
	}

	user, err := h.uc.CreateUser(context.Background(), &request)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	resp := e.HttpResp[e.UserCommonInfo]{Message: "Success", Result: user.Info}
	c.JSON(http.StatusOK, resp)
}
