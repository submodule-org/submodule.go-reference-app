package logger

import (
	"reference/env"

	"github.com/submodule-org/submodule.go"
	"go.uber.org/zap"
)

var LoggerMod = submodule.Make[*zap.Logger](func(env env.AppEnv) (z *zap.Logger, e error) {
	if env.IsProd() {
		return zap.NewProduction()
	}

	return zap.NewDevelopment()
}, env.EnvMod)
