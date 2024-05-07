package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/submodule-org/submodule.go"
)

func TestEnv(t *testing.T) {

	t.Run("default env", func(t *testing.T) {
		s := submodule.CreateScope()
		defer s.Dispose()

		env := EnvMod.ResolveWith(s)
		assert.True(t, env.IsDev())
		assert.False(t, env.IsProd())
		assert.False(t, env.IsTest())
	})

	t.Run("can be changed with APP_ENV", func(t *testing.T) {
		p := os.Getenv("APP_ENV")
		os.Setenv("APP_ENV", "prod")
		defer os.Setenv("APP_ENV", p)

		s := submodule.CreateScope()
		defer s.Dispose()

		env := EnvMod.ResolveWith(s)

		assert.True(t, env.IsProd())
		assert.False(t, env.IsDev())
		assert.False(t, env.IsTest())
	})
}
