package commands

import (
	"pock/internal/helpers"

	"github.com/spf13/cobra"
)

// NewInstallCommand creates the install command
func NewInstallCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install <name>",
		Short: "Install a command directly from the marketplace",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			helpers.PrintFeatureNotImplemented("Use 'pock import' to import commands from a URL or file.")
			return nil
		},
	}

	cmd.Flags().BoolP("force", "f", false, "Overwrite if command already exists")

	return cmd
}
