package download

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadFile(url string, dest string) error {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to download file from URL: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("URL returned HTTP %d status: %s", resp.StatusCode, resp.Body)
	}
	out, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("Failed to create file: %w", err)
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("Failed to write file: %w", err)
	}
	return nil
}
