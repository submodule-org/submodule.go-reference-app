package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/submodule-org/submodule.go"
)

func TestConfig(t *testing.T) {

	t.Run("default config", func(t *testing.T) {
		s := submodule.CreateScope()
		defer s.Dispose()

		config := ConfigMod.ResolveWith(s)

		assert.Equal(t, 8080, config.ServerPort)
		assert.Equal(t, "info", config.LogLevel)
	})

	t.Run("can change with env variable", func(t *testing.T) {
		s := submodule.CreateScope()
		defer s.Dispose()

		os.Setenv("SERVER_PORT", "1234")
		os.Setenv("LOG_LEVEL", "debug")

		config := ConfigMod.ResolveWith(s)
		assert.Equal(t, 1234, config.ServerPort)
		assert.Equal(t, "debug", config.LogLevel)
	})

	t.Run("invalid port may cause error", func(t *testing.T) {
		s := submodule.CreateScope()
		defer s.Dispose()

		os.Setenv("SERVER_PORT", "abcd")
		os.Setenv("LOG_LEVEL", "debug")

		_, e := ConfigMod.SafeResolveWith(s)
		assert.Error(t, e)

	})
}
