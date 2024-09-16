package commands

import (
	"cli/utility/download"
	"cli/utility/filehandler"
	"fmt"

	"cli/outputs"

	"github.com/spf13/cobra"
)

func GetCopyCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "copy",
		Short: "Copy files to a destination",
		Long:  "Copy a file or directory to a destination, maintains permissions",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return fmt.Errorf("Not enough arguments, expected 2 but got %d", len(args))
			}
			source := args[0]
			destination := args[1]
			err := filehandler.Copy(source, destination)
			if err != nil {
				return fmt.Errorf("Failed to copy: %w", err)
			}
			outputs.PrintStandardMessage("COPY", "Successfully copied %s: %s", source, destination)
			return nil
		},
	}
}

func GetDownloadCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "download",
		Short: "Download file from URL",
		Long:  "Download a file from a URL and save to a given location",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return fmt.Errorf("Not enough arguments, expected 2 but got %d", len(args))
			}
			url := args[0]
			destination := args[1]
			err := download.DownloadFile(url, destination)
			if err != nil {
				return fmt.Errorf("Failed to download: %w", err)
			}
			outputs.PrintStandardMessage("DOWNLOAD", "Successfully downloaded to: %s", destination)
			return nil
		},
	}
}
