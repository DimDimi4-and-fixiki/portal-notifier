package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	e "notify/internal/entity"
)

var info e.UserCommonInfo
var cred e.UserCred

func bindUserData(c *gin.Context) {
	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&cred); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func (h *Handler) RegisterUser(c *gin.Context) {
	bindUserData(c)
	res, err := h.uc.CreateUserWithDetails(info, cred)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"result": res})
}
