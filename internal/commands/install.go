package commands

import (
	"fmt"
	"pock/internal/utils"

	"github.com/spf13/cobra"
)

// NewInstallCommand creates the install command
func NewInstallCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install <name>",
		Short: "Install a command directly from the marketplace",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// This is a placeholder for marketplace integration
			fmt.Printf("%s\n", utils.Yellow("Marketplace integration not yet implemented."))
			fmt.Printf("%s Use 'pock import' to import commands from a URL or file.\n", utils.Blue("ℹ"))
			return nil
		},
	}

	cmd.Flags().BoolP("force", "f", false, "Overwrite if command already exists")

	return cmd
}
