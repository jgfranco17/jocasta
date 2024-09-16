package commands

import (
	"cli/internal"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertFileExists(t *testing.T, filePath string) {
	t.Helper()
	_, err := os.Stat(filePath)
	assert.NoError(t, err)
}

func TestCopyCommandHelp(t *testing.T) {
	output := internal.ExecuteTestCommand(GetCopyCommand, "--help")
	assert.NoError(t, output.Error)
	assert.Contains(t, output.ShellOutput, "copy [flags]", "Did not find usage guide for copy command")
}

func TestCopyCommandNotEnoughArgs(t *testing.T) {
	output := internal.ExecuteTestCommand(GetCopyCommand, "/some/path")
	assert.ErrorContains(t, output.Error, "expected 2 but got 1")
}

func TestCopyCommandCopySuccess(t *testing.T) {
	tempDir := internal.CreateTempTestingDir(t)
	defer os.RemoveAll(tempDir)
	sampleFile := filepath.Join(tempDir, "copied_testfile.txt")
	output := internal.ExecuteTestCommand(GetCopyCommand, "./resources/sample.txt", sampleFile)
	assert.NoError(t, output.Error)
	assertFileExists(t, sampleFile)
}

func TestDownloadFileSuccess(t *testing.T) {
	// Set up a mock server that returns a success response
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "mock content")
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()
	tempDir := t.TempDir()
	destFile := tempDir + "/testfile.txt"

	output := internal.ExecuteTestCommand(GetDownloadCommand, server.URL, destFile)
	assert.NoError(t, output.Error, "Expected no error when downloading file")
	_, err := os.Stat(destFile)
	assert.NoError(t, err, "Expected the file to be created")
	assertFileExists(t, destFile)
}
