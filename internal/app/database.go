package app

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"notify/internal/config"
	"os"
	"time"
)

func pgDSN(cfg config.DbConfig) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s port=%s dbname=%s sslmode=disable TimeZone=Europe/Moscow",
		cfg.Host, cfg.User, cfg.Password, cfg.Port, cfg.Name,
	)
}

func pgConfig(cfg config.DbConfig) postgres.Config {
	return postgres.Config{DSN: pgDSN(cfg), PreferSimpleProtocol: true}
}

func (a *App) newPgConnect(cfg config.DbConfig) (*gorm.DB, error) {
	gormCfg := gorm.Config{}
	if cfg.LogSql {
		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             0 * time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info,     // Log level
				IgnoreRecordNotFoundError: false,           // Ignore ErrRecordNotFound error for logger
				ParameterizedQueries:      true,            // Don't include params in the SQL log
				Colorful:                  true,            // Disable color
			},
		)
		gormCfg.Logger = newLogger
	}

	db, err := gorm.Open(postgres.New(pgConfig(cfg)), &gormCfg)
	return db, err
}

func (a *App) createUUIDPgExtension(conn *gorm.DB) error {
	res := conn.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	return res.Error
}
