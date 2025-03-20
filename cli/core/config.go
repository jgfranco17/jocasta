package core

import (
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

const (
	projectAppDirectory string = ".jocasta"
)

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || !os.IsNotExist(err) // Returns true if file exists
}

func FindAppDirectory() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Failed to identify home directory: %w", err)
	}
	appDir := filepath.Join(homeDir, projectAppDirectory)
	if !fileExists(appDir) {
		err := os.MkdirAll(appDir, 0755)
		if err != nil {
			return "", fmt.Errorf("Failed to create app directory: %w", err)
		}
		log.Debugf("App directory not found, creating new: %s", appDir)
	}
	return appDir, nil
}
