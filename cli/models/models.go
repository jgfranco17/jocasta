package models

import (
	"github.com/spf13/cobra"
)

type JocastaCommandFunction func() *cobra.Command

type CommandRunner func(cmd *cobra.Command, args []string)
