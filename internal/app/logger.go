package app

import (
	"log"
	"notify/pkg/logger"
)

func (a *App) initLogger() {
	l, err := logger.NewLogger()
	if err != nil {
		log.Fatal(err)
	}

	a.logger = l
}
