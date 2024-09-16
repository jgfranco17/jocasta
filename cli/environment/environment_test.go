package environment

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfigValueSet(t *testing.T) {
	t.Setenv("JOCASTA_LOG_LEVEL", "INFO")
	cfg, err := LoadConfigs()
	assert.NoError(t, err)
	assert.Equal(t, "INFO", cfg.LogLevel())
}
