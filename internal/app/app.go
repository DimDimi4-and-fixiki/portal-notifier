package app

import (
	"context"
	"errors"
	courierc "github.com/trycourier/courier-go/v3"
	"go.uber.org/zap/zapcore"
	"gopkg.in/tylerb/graceful.v1"
	"gorm.io/gorm"
	"notify/internal/config"
	e "notify/internal/entity"
	"notify/internal/handler/http"
	"sync"
)

type Logger interface {
	Debug(string, ...zapcore.Field)
	Info(string, ...zapcore.Field)
	Error(string, ...zapcore.Field)
	Fatal(string, ...zapcore.Field)
}

type HttpServer interface {
	GetGracefulServer() *graceful.Server
	Start() error
	RegisterRoutes(r *http.Router)
}

type CourierClient interface {
	SendTextMessage(context.Context, e.EmailMsg) (*courierc.SendMessageResponse, error)
}

type App struct {
	cfg config.Config

	c     *Container
	cOnce *sync.Once

	pg            *gorm.DB
	server        HttpServer
	courierClient CourierClient

	logger Logger
}

var a *App

func NewApp() (*App, error) {
	cfg := config.Get()
	app := &App{
		cOnce: &sync.Once{},
		//hcOnce: &sync.Once{},
		cfg: cfg,
	}

	//goro:init logger
	app.initLogger()

	//goro:init healthChecker
	//app.initHealthChecker()

	pgConn, err := app.newPgConnect(cfg.Db)
	if err != nil {
		return nil, err
	}

	err = app.createUUIDPgExtension(pgConn)
	if err != nil {
		return nil, err
	}

	app.pg = pgConn
	app.server = app.newGracefulServer()
	app.courierClient = app.newCourierClient(cfg.Courier)

	app.c = NewContainer(app.pg, app.server.GetGracefulServer(), app.courierClient)

	return app, nil
}

func SetGlobalApp(app *App) {
	a = app
}

func GetGlobalApp() (*App, error) {
	if a == nil {
		return nil, errors.New("global app is not initialized")
	}

	return a, nil
}
