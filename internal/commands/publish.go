package commands

import (
	"fmt"
	"pock/internal/utils"

	"github.com/spf13/cobra"
)

// NewPublishCommand creates the publish command
func NewPublishCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "publish <name>",
		Short: "Publish a command to the marketplace",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// This is a placeholder for marketplace integration
			fmt.Printf("%s\n", utils.Yellow("Marketplace integration not yet implemented."))
			fmt.Printf("%s Use 'pock export' to export commands to a file.\n", utils.Blue("ℹ"))
			return nil
		},
	}

	return cmd
}
