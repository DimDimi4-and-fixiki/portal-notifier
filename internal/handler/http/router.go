package http

import (
	ginpprof "github.com/gin-contrib/pprof"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	cfg "notify/internal/config"
	"notify/pkg/logger"
)

type Router struct {
	router *gin.Engine
}

func NewRouter() *Router {
	router := gin.New()
	router.Use(gin.Recovery())
	return &Router{router: router}
}

func (r *Router) WithDebug() *Router {
	ginpprof.Register(r.router)
	return r
}

func (r *Router) WithLogging() (*Router, error) {
	log := logger.NewCustomLogger()
	r.router.Use(ginzap.Ginzap(
		log,
		cfg.Get().Server.LogTimeFormat, true))
	return r, nil
}

func (r *Router) Router() *gin.Engine {
	return r.router
}
