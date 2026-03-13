package commands

import (
	"fmt"
	"pock/internal/utils"

	"github.com/spf13/cobra"
)

// NewBrowseCommand creates the browse command
func NewBrowseCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "browse",
		Short: "Browse commands in the marketplace",
		RunE: func(cmd *cobra.Command, args []string) error {
			// This is a placeholder for marketplace integration
			fmt.Printf("%s\n", utils.Yellow("Marketplace integration not yet implemented."))
			fmt.Printf("%s Use 'pock import' to import commands from a URL or file.\n", utils.Blue("ℹ"))
			return nil
		},
	}

	return cmd
}
