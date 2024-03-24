package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	e "notify/internal/entity"
	http_ "notify/internal/handler/http"
)

type respSendProposal struct{}

// SendProposal godoc
// @Summary Send email with proposal to user and owner
// @Description Send email with proposal text to user and owner
// @Tags v1
// @Accept application/json
// @Produce application/json
// @Param data body e.ReqWithApiToken[e.ProjectProposal] true "Project proposal data"
// @Success 200 {object} http_.Resp[respSendProposal]
// @Failure 400 {object} http_.RespErr
// @Failure 401 {object} http_.RespErr
// @Router /v1/send_proposal/ [post]
func (h *Handler) SendProposal(c *gin.Context) {
	var request e.ReqWithApiToken[e.ProjectProposal]
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, http_.RespValidationErr(err))
		h.logger.Info("request validation error", zap.Error(err))
		return
	}

	ctx := context.Background()
	auth, err := h.uc.AuthUserApiToken(ctx, request.Auth.Login, request.Auth.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, http_.RespAuthErr(err))
		return
	}

	if !auth {
		c.JSON(http.StatusUnauthorized, http_.RespUnauthorizedErr())
		return
	}

	err = h.uc.SendProposal(ctx, request.Data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, http_.RespInternalErr(err))
		return
	}
	resp := http_.Resp[respSendProposal]{Message: "Success", Result: respSendProposal{}}
	c.JSON(http.StatusOK, resp)
}
