package cmd

import (
	"pock/internal/helpers"

	"github.com/spf13/cobra"
)

// NewPublishCommand creates the publish command
func NewPublishCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "publish <name>",
		Short: "Publish a command to the marketplace",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			helpers.PrintFeatureNotImplemented("Use 'pock export' to export commands to a file.")
			return nil
		},
	}

	return cmd
}
