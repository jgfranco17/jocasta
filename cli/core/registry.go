package core

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	verbosity int
)

type CommandRegistry struct {
	rootCmd         *cobra.Command
	verbosity       int
	configDirectory string
}

// NewCommandRegistry creates a new instance of CommandRegistry
func NewCommandRegistry(name string, description string, version string) *CommandRegistry {
	root := &cobra.Command{
		Use:     name,
		Version: version,
		Short:   description,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			verbosity, _ := cmd.Flags().GetCount("verbose")
			switch verbosity {
			case 1:
				logrus.SetLevel(logrus.InfoLevel)
			case 2:
				logrus.SetLevel(logrus.DebugLevel)
			case 3:
				logrus.SetLevel(logrus.TraceLevel)
			default:
				logrus.SetLevel(logrus.WarnLevel)
			}
		},
	}
	newRegistry := &CommandRegistry{
		rootCmd:         root,
		configDirectory: "",
	}
	root.PersistentFlags().CountVarP(&newRegistry.verbosity, "verbose", "v", "Increase verbosity (-v or -vv)")
	root.Flags().BoolP("version", "V", false, "Print the version number of Jocasta")
	return newRegistry
}

// RegisterCommand registers a new command with the CommandRegistry
func (cr *CommandRegistry) RegisterCommands(commands []*cobra.Command) {
	for _, cmd := range commands {
		cr.rootCmd.AddCommand(cmd)
	}
}

// SetConfigDirectory sets the directory where configuration files will be stored
func (cr *CommandRegistry) SetConfigDirectory() error {
	appDirectory, err := FindAppDirectory()
	if err != nil {
		return err
	}
	cr.configDirectory = appDirectory
	return nil
}

// Execute executes the root command
func (cr *CommandRegistry) Execute() error {
	return cr.rootCmd.Execute()
}
