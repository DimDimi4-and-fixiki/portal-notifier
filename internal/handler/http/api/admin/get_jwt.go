package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var body struct {
	Login    string
	Password string
}

func (h *Handler) GetJWT(c *gin.Context) {
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}
