package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	pathFromEnv, isEnvSet := os.LookupEnv(EnvConfigPath)
	if isEnvSet && pathFromEnv != "" {
		os.Unsetenv(EnvConfigPath)
	}
	code := m.Run()
	os.Setenv(EnvConfigPath, pathFromEnv)
	os.Exit(code)
}

func TestGetConfigPathSuccess(t *testing.T) {
	mockConfigPath := "/some/path"
	t.Setenv(EnvConfigPath, mockConfigPath)

	assert.Equal(t, mockConfigPath, GetConfigPathFromEnv())
}

func TestGetConfigPathEnvSetEmpty(t *testing.T) {
	assert.Emptyf(t, GetConfigPathFromEnv(), "Variable %s was not empty", EnvConfigPath)
}
