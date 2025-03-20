package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"cli/core"
	"cli/logging"
)

const (
	projectName        = "jocasta"
	projectDescription = "Jocasta: Track your tasks from the command line."
)

var (
	version string = "0.0.0-dev"
	command *core.CommandRegistry
)

func init() {
	// Configure logging
	log.SetReportCaller(true)
	log.SetFormatter(&logging.CustomFormatter{})

	// Register commands to the command registry
	commandsList := []*cobra.Command{
		core.GetListCommand(),
	}
	command = core.NewCommandRegistry(projectName, projectDescription, version)
	command.RegisterCommands(commandsList)
}

func main() {
	err := command.SetConfigDirectory()
	if err != nil {
		log.Fatalf("App config directory could not be set: %v", err)
	}
	err = command.Execute()
	if err != nil {
		log.Error(err.Error())
	}
}
