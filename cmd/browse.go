package cmd

import (
	"pock/internal/helpers"

	"github.com/spf13/cobra"
)

// NewBrowseCommand creates the browse command
func NewBrowseCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "browse",
		Short: "Browse commands in the marketplace",
		RunE: func(cmd *cobra.Command, args []string) error {
			helpers.PrintFeatureNotImplemented("Use 'pock import' to import commands from a URL or file.")
			return nil
		},
	}

	return cmd
}
