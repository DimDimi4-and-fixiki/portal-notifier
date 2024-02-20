package admin

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	e "notify/internal/entity"
	http_ "notify/internal/handler/http"
)

// @Summary get user info
// @Description Return user by user id or user login
// @Tags Admin
// @Accept application/json
// @Produce application/json
// @Param user body e.GetUserReq true "no_comm"
// @Success 200 {object} http_.Resp[e.UserCommonInfo]
// @Failure 400 {object} http_.RespErr
// @Failure 200 {object} http_.RespErr
// @Router /admin/get_user/ [post]
func (h *Handler) getUser(c *gin.Context) {
	var input e.GetUserReq
	err := c.BindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, http_.RespValidationErr(err))
		return
	}
	h.logger.Info("Parsed input")
	user, err := h.uc.GetUser(context.Background(), input)
	if err != nil {
		c.JSON(http.StatusOK, http_.RespInternalErr(err))
		return
	}

	resp := http_.Resp[e.UserCommonInfo]{Message: "Success", Result: user.Info}
	c.JSON(http.StatusOK, resp)
}
