package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	e "notify/internal/entity"
	http_ "notify/internal/handler/http"
)

type respDecodeJWT struct {
	UserLogin string     `json:"user_login"`
	UserRole  e.UserRole `json:"user_role"`
}

type req struct {
	Token string `json:"token"`
}

// DecodeJWT godoc
// @Summary View what's inside your JWT token
// @Description Decode JWT token
// @Tags v1
// @Accept application/json
// @Produce application/json
// @Param req body req true "JWT token"
// @Success 200 {object} http_.Resp[respDecodeJWT]
// @Failure 400 {object} http_.RespErr
// @Router /v1/decode_jwt/ [post]
func (h *Handler) DecodeJWT(c *gin.Context) {
	var request req
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, http_.RespValidationErr(err))
		h.logger.Info("request validation error", zap.Error(err))
		return
	}

	ctx := context.Background()
	user, err := h.uc.DecodeJWTToUser(ctx, request.Token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, http_.Resp[respDecodeJWT]{
		Message: "Success",
		Result:  respDecodeJWT{UserLogin: user.UserLogin, UserRole: user.UserRole},
	},
	)

}
