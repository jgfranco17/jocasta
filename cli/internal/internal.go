package internal

import (
	"bytes"
	"cli/models"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type CliRunResult struct {
	ShellOutput string
	Error       error
}

// Helper function to simulate CLI execution
func ExecuteTestCommand(cmdGetter models.JocastaCommandFunction, args ...string) CliRunResult {
	buf := new(bytes.Buffer)
	cmd := cmdGetter()
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs(args)

	_, err := cmd.ExecuteC()
	return CliRunResult{
		ShellOutput: buf.String(),
		Error:       err,
	}
}

func CreateTempTestingDir(t *testing.T) string {
	t.Helper()
	tempDir, err := os.MkdirTemp("", "jocasta-test-*")
	assert.NoError(t, err)
	return tempDir
}

func CreateTempFile(t *testing.T, directory string, fileName string) *os.File {
	t.Helper()
	tempFile, err := os.CreateTemp(directory, fileName)
	assert.NoError(t, err)
	return tempFile
}
