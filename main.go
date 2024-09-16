package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"cli/commands"
	"cli/environment"
	"cli/logging"
)

const (
	projectName        string = "jocasta"
	projectDescription string = "File management, simplified.\nDeveloped by Joaquin Franco."
)

func init() {
	log.SetReportCaller(true)
	log.SetFormatter(&logging.CustomFormatter{})
}

func main() {
	cfg, err := environment.LoadConfigs()
	logging.ConfigureLogger(cfg.LogLevel())
	commandsList := []*cobra.Command{
		commands.GetCopyCommand(),
		commands.GetDownloadCommand(),
	}
	command := commands.NewCommandRegistry(projectName, projectDescription)
	command.RegisterCommands(commandsList)

	err = command.Execute()
	if err != nil {
		log.Error(err.Error())
	}
}
