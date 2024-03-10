package admin

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"notify/internal/config"
	e "notify/internal/entity"
	http_ "notify/internal/handler/http"
)

// RegisterService godoc
// @Summary Register new service
// @Description Register new user with user details
// @Tags Admin
// @Accept application/json
// @Produce application/json
// @Param user body e.CreateServiceReq true "Service details"
// @Success 200 {object} http_.Resp[e.ServiceUser]
// @Failure 400 {object} http_.RespErr
// @Router /admin/register_service/ [post]
func (h *Handler) RegisterService(c *gin.Context) {
	var request e.CreateServiceReq
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, http_.RespValidationErr(err))
		h.logger.Info("request validation error", zap.Error(err))
		return
	}

	reqUser, err := request.ToUser(config.Get().Crypt)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	user, err := h.uc.CreateService(context.Background(), reqUser)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	resp := http_.Resp[e.ServiceUser]{Message: "Success", Result: user}
	c.JSON(http.StatusOK, resp)
}
