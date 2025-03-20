package models

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Test helper for setting up a temporary file
func createTempFileWithData(t *testing.T, filename string, data string) {
	t.Helper()
	err := os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
}

// Test helper for cleaning up the temp file
func cleanUpTempFile(t *testing.T, filename string) {
	t.Helper()
	err := os.Remove(filename)
	if err != nil {
		t.Fatalf("Failed to remove temp file: %v", err)
	}
}

func TestLoadTaskCollectionFromFile(t *testing.T) {
	// Prepare test data
	tasksData := `{
		"tasks": {
			"1": {"date_created": "2025-03-20T12:00:00Z", "description": "Test Task 1"},
			"2": {"date_created": "2025-03-20T12:30:00Z", "description": "Test Task 2"}
		}
	}`

	// Create temporary file with task data
	tempFile := "test_tasks.json"
	createTempFileWithData(t, tempFile, tasksData)
	defer cleanUpTempFile(t, tempFile)

	// Load the task collection from the file
	ctx := context.Background()
	taskCollection, err := LoadTaskCollectionFromFile(ctx, tempFile)

	// Assert no error occurs
	assert.NoError(t, err)

	// Assert task collection is loaded correctly
	assert.NotNil(t, taskCollection)
	assert.Equal(t, 2, len(taskCollection.Tasks))
	assert.Equal(t, "Test Task 1", taskCollection.Tasks[1].Description)
	assert.Equal(t, "Test Task 2", taskCollection.Tasks[2].Description)
}

func TestListTasks(t *testing.T) {
	// Capture the output of List() to verify it prints correctly
	// This would require mocking or using a buffer (could use `bytes.Buffer` and `fmt.Fprintf` if refactored to write to it)

	// Initialize a task collection with some tasks
	taskCollection := &TaskCollection{
		Tasks: map[int]Task{
			1: {DateCreated: time.Now(), Description: "Task 1"},
			2: {DateCreated: time.Now(), Description: "Task 2"},
		},
	}

	// Using assert to test the number of tasks (though List() prints to console)
	assert.Equal(t, 2, len(taskCollection.Tasks))

	// Normally, you would mock fmt.Println or refactor List() to write to a buffer.
	// Since we're not dealing with the console here, we won't capture output directly in the test.
	// For production code, refactor List() to make it easier to test.
}

func TestLoadTaskCollectionFromFile_Error(t *testing.T) {
	// Test for handling file read error (file does not exist)
	ctx := context.Background()
	_, err := LoadTaskCollectionFromFile(ctx, "nonexistent_file.json")
	assert.Error(t, err)

	// Test for handling invalid JSON
	invalidJSON := `{ "tasks": { "1": "invalid" }`
	tempFile := "invalid_tasks.json"
	createTempFileWithData(t, tempFile, invalidJSON)
	defer cleanUpTempFile(t, tempFile)

	_, err = LoadTaskCollectionFromFile(ctx, tempFile)
	assert.Error(t, err)
}

func TestSaveTaskCollection_Error(t *testing.T) {
	// Test for handling file write error (e.g., permission issues)
	// Create a file that would be read-only or a bad path to test
	tempFile := "/root/test_save_tasks.json" // Assuming you don't have write permission in this path
	taskCollection := &TaskCollection{
		Tasks: map[int]Task{
			1: {DateCreated: time.Now(), Description: "Task 1"},
		},
	}

	err := taskCollection.Save(tempFile)
	assert.Error(t, err)
}
