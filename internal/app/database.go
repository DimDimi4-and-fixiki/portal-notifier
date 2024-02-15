package app

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"notify/internal/config"
	"notify/pkg/logger"
)

func pgDSN(cfg config.DbConfig) string {
	logger.NewSugar().Info(fmt.Sprintf(
		"host=%s user=%s password=%s port=%s dbname=%s sslmode=disable TimeZone=Europe/Moscow",
		cfg.Host, cfg.User, cfg.Password, cfg.Port, cfg.Name,
	))
	return fmt.Sprintf(
		"host=%s user=%s password=%s port=%s dbname=%s sslmode=disable TimeZone=Europe/Moscow",
		cfg.Host, cfg.User, cfg.Password, cfg.Port, cfg.Name,
	)
}

func pgConfig(cfg config.DbConfig) postgres.Config {
	return postgres.Config{DSN: pgDSN(cfg), PreferSimpleProtocol: true}
}

func (a *App) newPgConnect(cfg config.DbConfig) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(pgConfig(cfg)), &gorm.Config{})
	return db, err
}

func (a *App) createUUIDPgExtension(conn *gorm.DB) error {
	res := conn.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	return res.Error
}
