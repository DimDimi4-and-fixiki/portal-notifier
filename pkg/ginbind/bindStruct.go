package ginbind

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"notify/pkg/logger"
)

var log = logger.NewSugar()

type BindJsonError struct {
	err error
}

func (e BindJsonError) Error() string {
	return fmt.Sprintf("Error while Binding Request to JSON: %s", e.err.Error())
}

// BindJSONWithLog  tries to bind request to JSON struct
// if error occurred, returns it and logs
func BindJSONWithLog(c *gin.Context, model any) error {
	err := c.BindJSON(model)
	if err != nil {
		log.Error(zap.Error(err))
	}
	return err
}
