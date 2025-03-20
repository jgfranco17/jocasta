package core

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	failFast bool
)

func GetListCommand() *cobra.Command {
	var count int
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Show all tasks",
		Long:  "Get the full list of all your current tasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			homeDir, err := os.UserHomeDir()
			if err != nil {
				return fmt.Errorf("Failed to identify home directory: %w", err)
			}

			filepath := filepath.Join(homeDir, projectAppDirectory, "tasks.json")
			err = createFile(filepath, "{}")
			if err != nil {
				return fmt.Errorf("Failed to create tasks file: %w", err)
			}
			taskCollection, err := LoadTaskCollectionFromFile(ctx, filepath)
			if err != nil {
				return fmt.Errorf("Error loading task collection: %w", err)
			}
			taskCollection.List()
			return nil
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	cmd.Flags().IntVarP(&count, "limit", "c", 1, "Number of ping requests, default is 1")
	return cmd
}

func createFile(path string, content string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close() // Ensure the file is closed
	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}
