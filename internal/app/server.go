package app

import (
	"go.uber.org/zap"
	"notify/internal/handler/http"
	admin "notify/internal/handler/http/api/admin"
	v1 "notify/internal/handler/http/api/v1"
)

func (a *App) newGracefulServer() IHttpServer {
	return http.NewGracefulServer(a.cfg.Server)
}

func (a *App) StartHTTPServer() {
	router, err := http.NewRouter().WithLogging()
	a.server.RegisterRoutes(router)
	if err != nil {
		a.logger.Fatal("Fail to start http server:", zap.Error(err))
	}

	v1.NewHandler(a.c.GetUseCase(), a.logger, router.Router()).WithRoutes()
	admin.NewHandler(a.c.GetUseCase(), a.logger, router.Router()).WithRoutes()

	err = a.server.Start()
	if err != nil {
		a.logger.Fatal("Fail to start http server:", zap.Error(err))
	}

}
