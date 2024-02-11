package http

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/tylerb/graceful.v1"
	"net/http"
	"notify/internal/config"
	"time"
)

type Server struct {
	srv *graceful.Server
}

// Create Server with graceful shutdown timeout
func newGracefulServer(cfg config.HttpServerConfig) *graceful.Server {
	return &graceful.Server{
		Timeout: time.Second * time.Duration(cfg.GracefulTimeoutSec),
		Server: &http.Server{
			Addr: fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
			//WriteTimeout: cfg.WriteTimeout,
		},
	}
}

func NewGracefulServer(cfg config.HttpServerConfig) *Server {
	return &Server{
		srv: newGracefulServer(cfg),
	}
}

func (s *Server) RegisterRoutes(r *Router) {
	s.srv.Handler = r.router
}

func (s *Server) GetGracefulServer() *graceful.Server {
	return s.srv
}

func (s *Server) Start() error {
	if s.srv.Handler == nil {
		return fmt.Errorf("no routes registered on server handler :(")
	}

	err := s.srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (s *Server) Stop() error {
	return s.srv.Shutdown(context.Background())
}
