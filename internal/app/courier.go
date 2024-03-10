package app

import (
	"notify/internal/config"
	"notify/internal/repo/clients/courier"
)

func (a *App) newCourierClient(cfg config.CourierConfig) *courier.Client {
	return courier.NewClient(cfg)
}
