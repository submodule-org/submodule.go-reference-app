package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"reference/common"
	"reference/logger"
	"reference/server"

	"github.com/submodule-org/submodule.go"
	"go.uber.org/zap"
)

func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	logger := logger.LoggerMod.Resolve()
	submodule.AppendGlobalMiddleware(
		submodule.WithScopeResolve(func(r common.Registry) common.Registry {
			o := r.Handler
			r.Handler = func(w http.ResponseWriter, r *http.Request) {
				logger.Info("Request", zap.String("path", r.URL.Path))
				defer logger.Info("Response", zap.String("path", r.URL.Path))

				o(w, r)
			}

			return r
		}),
	)

	server := server.ServerMod.Resolve()

	go func() {
		e := server.Start()
		if e != nil && !errors.Is(e, http.ErrServerClosed) {
			logger.Fatal("Failed to start server", zap.Error(e))
		}
	}()

	sig := <-signalChan
	logger.Info("Received signal", zap.String("signal", sig.String()))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Stop(ctx); err != nil {
		logger.Error("Failed to gracefully shutdown server", zap.Error(err))
	}

	logger.Info("Server stopped")
}
