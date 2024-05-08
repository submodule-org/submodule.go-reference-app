package common

import (
	"fmt"
	"net/http"
	"reference/logger"
	"strings"

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
	path = strings.TrimSpace(path)
	path = strings.TrimPrefix(path, "/")
	path = strings.TrimSuffix(path, "/")
	path = fmt.Sprintf("/%s/", path)

	rs := submodule.Group[Registry](registries...)

	return submodule.Make[Mux](func(handlers []Registry, logger *zap.Logger) Mux {
		mux := http.NewServeMux()

		for _, registry := range handlers {
			mux.HandleFunc(registry.Path, registry.Handler)
		}

		return Mux{
			Path: path,
			Mux:  mux,
		}
	}, rs, logger.LoggerMod)
}
