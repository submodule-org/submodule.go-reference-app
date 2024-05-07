package env

import (
	"os"

	"github.com/submodule-org/submodule.go"
)

type appEnv struct {
	env string
}

type AppEnv interface {
	IsProd() bool
	IsDev() bool
	IsTest() bool
}

func (a *appEnv) IsProd() bool {
	return a.env == "production" || a.env == "prod"
}

func (a *appEnv) IsDev() bool {
	return a.env == "" || a.env == "development" || a.env == "dev"
}

func (a *appEnv) IsTest() bool {
	return a.env == "test"
}

var EnvMod = submodule.Make[AppEnv](func() AppEnv {
	env := os.Getenv("APP_ENV")
	return &appEnv{
		env: env,
	}
})
