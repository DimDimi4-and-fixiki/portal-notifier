package admin

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	e "notify/internal/entity"
	http_ "notify/internal/handler/http"
)

// RegisterUser godoc
// @Summary Register new user
// @Description Register new user with user details
// @Tags Admin
// @Accept application/json
// @Produce application/json
// @Param user body e.User true "User details"
// @Success 200 {object} http_.Resp[e.UserCommonInfo]
// @Failure 400 {object} http_.RespErr
// @Router /admin/register_user/ [post]
func (h *Handler) RegisterUser(c *gin.Context) {
	var request e.User
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, http_.RespValidationErr(err))
		h.logger.Info("request validation error", zap.Error(err))
		return
	}

	user, err := h.uc.CreateUser(context.Background(), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := http_.Resp[e.UserCommonInfo]{Message: "Success", Result: user.Info}
	c.JSON(http.StatusOK, resp)
}
