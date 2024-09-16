package download

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Mock server for testing
func mockServer(statusCode int, body string) *httptest.Server {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		io.WriteString(w, body)
	}
	return httptest.NewServer(http.HandlerFunc(handler))
}

func TestDownloadFileSuccess(t *testing.T) {
	// Set up a mock server that returns a success response
	server := mockServer(http.StatusOK, "mock content")
	defer server.Close()
	tempDir := t.TempDir() // Create a temporary directory
	destFile := tempDir + "/testfile.txt"

	err := DownloadFile(server.URL, destFile)
	assert.NoError(t, err, "Expected no error when downloading file")
	_, err = os.Stat(destFile)
	assert.NoError(t, err, "Expected the file to be created")
	content, err := os.ReadFile(destFile)
	assert.NoError(t, err, "Expected no error reading the file")
	assert.Equal(t, "mock content", string(content), "Expected file content to match the mock content")
}

func TestDownloadFileHttpError(t *testing.T) {
	server := mockServer(http.StatusNotFound, "Not Found")
	defer server.Close()
	tempDir := t.TempDir()
	destFile := tempDir + "/testfile.txt"

	err := DownloadFile(server.URL, destFile)
	assert.Error(t, err, "Expected an error due to HTTP 404 status")
	assert.Contains(t, err.Error(), "URL returned HTTP 404 status", "Expected error to mention 404 status")
	_, err = os.Stat(destFile)
	assert.True(t, os.IsNotExist(err), "Expected file not to be created on HTTP error")
}
