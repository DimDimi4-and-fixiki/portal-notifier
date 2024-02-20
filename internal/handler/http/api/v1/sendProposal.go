package v1

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	e "notify/internal/entity"
	http_ "notify/internal/handler/http"
)

// @Summary send proposal notification by email
// @Description Sends message via email to owner of project and confirmation to user's email
// @Tags v1
// @Accept application/json
// @Produce application/json
// @Param proposal body e.SendProposalReq true "no_comm"
// @Success 200 {object} http_.Resp[e.SendProposalResp]
// @Failure 400 {object} http_.RespErr
// @Failure 200 {object} http_.RespErr
// @Router /v1/send_proposal/ [post]
func (h *Handler) sendProposal(c *gin.Context) {
	var req e.SendProposalReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, http_.RespValidationErr(err))
		return
	}

	proposal := req.ToProjectProposal()
	res, errs := h.uc.SendProposal(context.TODO(), proposal, h.cfg.Courier.OwnerEmail)
	resp := e.NewSendProposalResp().FromProjectProposal(res)

	if len(errs) == 0 {
		resp := http_.Resp[e.SendProposalResp]{Message: "Success", Result: *resp}
		c.JSON(http.StatusOK, resp)
		return
	} else {
		errRes := errors.New("")
		for _, err := range errs {
			errRes = fmt.Errorf("%w %w", errRes, err)
		}
		c.JSON(http.StatusOK, http_.RespInternalErr(errRes))
	}
}
