package config

import (
	"os"
	"strconv"

	"reference/env"
	"reference/logger"

	"github.com/joho/godotenv"
	"github.com/submodule-org/submodule.go"
	"go.uber.org/zap"
)

type Config struct {
	ServerPort int
	LogLevel   string
}

var acceptedLogLevel = []string{"debug", "info", "warn", "error"}

func verifyLogLevel(logLevel string) bool {
	for _, l := range acceptedLogLevel {
		if l == logLevel {
			return true
		}
	}
	return false
}

func loadEnv(env env.AppEnv, logger *zap.Logger) {
	var configError error
	var e error
	if env.IsDev() {
		configError = godotenv.Load(".env.dev")
		if e != nil {
			logger.Info("Failed to load .env.dev file", zap.Error(configError))
		} else {
			logger.Info("Loaded .env.dev file")
		}
	}

	if env.IsTest() {
		configError = godotenv.Load(".env.test")
		if e != nil {
			logger.Info("Failed to load .env.test file", zap.Error(configError))
		} else {
			logger.Info("Loaded .env.test file")
		}
	}

	configError = godotenv.Load()
	if e != nil {
		logger.Info("Failed to load .env file", zap.Error(configError))
	} else {
		logger.Info("Loaded .env file")
	}
}

var ConfigMod = submodule.Make[Config](func(env env.AppEnv, logger *zap.Logger) (c Config, e error) {
	loadEnv(env, logger)

	serverPortStr := os.Getenv("SERVER_PORT")
	if serverPortStr == "" {
		serverPortStr = "8080"
	}

	serverPort, e := strconv.Atoi(serverPortStr)
	if e != nil {
		return
	}
	c.ServerPort = serverPort

	logLevel := os.Getenv("LOG_LEVEL")
	ok := verifyLogLevel(logLevel)
	if !ok {
		logLevel = "info"
	}
	c.LogLevel = logLevel

	return
}, env.EnvMod, logger.LoggerMod)
