package filehandler

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func copyFile(src string, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("Failed to open source file: %w", err)
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("Failed to create destination file: %w", err)
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("Error copying file: %w", err)
	}

	srcInfo, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("Error getting source file info: %w", err)
	}
	err = os.Chmod(dst, srcInfo.Mode())
	if err != nil {
		return fmt.Errorf("Error setting destination file permissions: %w", err)
	}

	return nil
}

func copyDirectory(srcDir string, dstDir string) error {
	srcInfo, err := os.Stat(srcDir)
	if err != nil {
		return fmt.Errorf("Failed to get directory info: %w", err)
	}
	err = os.MkdirAll(dstDir, srcInfo.Mode())
	if err != nil {
		return fmt.Errorf("Failed to create destination directory: %w", err)
	}
	entries, err := os.ReadDir(srcDir)
	if err != nil {
		return fmt.Errorf("Failed to read source directory: %w", err)
	}

	for _, entry := range entries {
		srcPath := filepath.Join(srcDir, entry.Name())
		dstPath := filepath.Join(dstDir, entry.Name())

		if entry.IsDir() {
			err = copyDirectory(srcPath, dstPath)
			if err != nil {
				return err
			}
		} else {
			err = copyFile(srcPath, dstPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Copy copies a file or directory from src to dst.
func Copy(src string, dst string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("Failed to get source info: %w", err)
	}

	if srcInfo.IsDir() {
		return copyDirectory(src, dst)
	}
	return copyFile(src, dst)
}
