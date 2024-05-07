package common

import (
	"net/http"
	"reference/logger"

	"github.com/submodule-org/submodule.go"
	"go.uber.org/zap"
)

type Registry struct {
	Path    string
	Handler func(http.ResponseWriter, *http.Request)
}

type Mux struct {
	Path string
	Mux  *http.ServeMux
}

func MakeMuxMod(path string, registries ...submodule.Retrievable) submodule.Submodule[Mux] {

	rs := submodule.Group[Registry](registries...)

	return submodule.Make[Mux](func(handlers []Registry, logger *zap.Logger) Mux {
		mux := http.NewServeMux()

		for _, registry := range handlers {
			logger.Info("Registering handler", zap.String("path", registry.Path))
			mux.HandleFunc(registry.Path, registry.Handler)
		}

		return Mux{
			Path: path,
			Mux:  mux,
		}
	}, rs, logger.LoggerMod)
}
