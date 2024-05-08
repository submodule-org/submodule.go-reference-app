package server

import (
	"context"
	"fmt"
	"net/http"

	"reference/config"
	"reference/env"
	rh "reference/http"
	"reference/logger"

	"github.com/submodule-org/submodule.go"
	"go.uber.org/zap"
)

type server struct {
	Config config.Config
	Logger *zap.Logger
	AppEnv env.AppEnv
	Routes []rh.Mux
}

type Server interface {
	Start() error
	Stop(context.Context) error
}

func (s *server) Start() error {
	if !s.AppEnv.IsProd() {
		s.Logger.Info("Starting server in development mode")
	}

	r := &http.ServeMux{}

	for _, route := range s.Routes {
		s.Logger.Info("Registering route", zap.String("path", route.Path), zap.Any("router", route.Mux))
		r.Handle(route.Path, route.Mux)
	}

	return http.ListenAndServe(fmt.Sprintf(":%d", s.Config.ServerPort), r)
}

func (s *server) Stop(ctx context.Context) error {
	s.Logger.Info("Stopping server")
	return nil
}

var ServerMod = submodule.Resolve[Server](&server{}, config.ConfigMod, env.EnvMod, logger.LoggerMod, routes)
