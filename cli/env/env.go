package env

import (
	"os"
	"path/filepath"
)

const (
	EnvConfigPath string = "JOCASTA_CONFIG_DIR"
)

func GetConfigPathFromEnv() string {
	pathFromEnv, isEnvSet := os.LookupEnv(EnvConfigPath)
	if !isEnvSet || pathFromEnv != "" {
		return pathFromEnv
	}
	return filepath.Join(os.Getenv("HOME"), ".jocasta")
}
