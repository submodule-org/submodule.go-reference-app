package server

import (
	"context"
	"fmt"
	"net/http"

	"reference/common"
	"reference/config"
	"reference/env"
	"reference/logger"
	"reference/todo_route"

	"github.com/submodule-org/submodule.go"
	"go.uber.org/zap"
)

type server struct {
	server *http.Server
	logger *zap.Logger
	appEnv env.AppEnv
}

type Server interface {
	Start() error
	Stop(context.Context) error
}

func (s *server) Start() error {
	if !s.appEnv.IsProd() {
		s.logger.Info("Starting server in development mode")
	}

	s.logger.Info("Starting server", zap.String("port", s.server.Addr))
	return s.server.ListenAndServe()
}

func (s *server) Stop(ctx context.Context) error {
	s.logger.Info("Stopping server")
	return s.server.Shutdown(ctx)
}

var muxes = submodule.Group[common.Mux](todo_route.TodoMuxMod)

var ServerMod = submodule.Make[Server](func(
	config config.Config,
	env env.AppEnv,
	logger *zap.Logger,
	muxes []common.Mux,
) (s Server, e error) {
	m := http.NewServeMux()
	for _, mux := range muxes {
		logger.Info("Registering handler", zap.String("path", mux.Path))
		m.Handle(mux.Path, mux.Mux)
	}

	s = &server{
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", config.ServerPort),
			Handler: m,
		},
		logger: logger,
		appEnv: env,
	}

	return
}, config.ConfigMod, env.EnvMod, logger.LoggerMod, muxes)
