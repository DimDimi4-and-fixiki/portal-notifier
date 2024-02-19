package app

import (
	"notify/internal/clients/courier"
	"notify/internal/config"
)

func (a *App) newCourierClient(cfg config.CourierConfig) *courier.Client {
	return courier.NewClient(cfg)
}
