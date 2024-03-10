package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	e "notify/internal/entity"
	http_ "notify/internal/handler/http"
)

type resp struct {
	User  e.UserCommonInfo `json:"user"`
	Token e.JWTString      `json:"jwt"`
}

// RegisterUser godoc
// @Summary Register new user with role person
// @Description Register new user with user details
// @Tags v1
// @Accept application/json
// @Produce application/json
// @Param user body e.ReqWithAuth[e.User] true "Auth data and User details"
// @Success 200 {object} http_.Resp[resp]
// @Failure 400 {object} http_.RespErr
// @Failure 401 {object} http_.RespErr
// @Router /v1/register_user/ [post]
func (h *Handler) RegisterUser(c *gin.Context) {
	var request e.ReqWithAuth[e.User]
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, http_.RespValidationErr(err))
		h.logger.Info("request validation error", zap.Error(err))
		return
	}

	ctx := context.Background()
	auth, err := h.uc.AuthUserApiToken(ctx, request.Auth.Login, request.Auth.Password)
	if err != nil {
		c.JSON(http.StatusOK, http_.RespAuthErr(err))
		return
	}

	if !auth {
		c.JSON(http.StatusUnauthorized, http_.RespUnauthorizedErr())
		return
	}

	user, err := h.uc.CreateUser(ctx, &request.Data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	token, err := h.uc.JWTFromUser(ctx, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := http_.Resp[resp]{Message: "Success", Result: resp{User: user.Info, Token: *token}}
	c.JSON(http.StatusOK, resp)
}
