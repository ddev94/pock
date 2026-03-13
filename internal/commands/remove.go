package commands

import (
	"fmt"
	"pock/internal/storage"
	"pock/internal/utils"

	"github.com/spf13/cobra"
)

// NewRemoveCommand creates the remove command
func NewRemoveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove <name>",
		Short: "Remove a saved command",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]

			// Find the command
			savedCommand, err := storage.GetSavedCommandByName(name)
			if err != nil {
				return fmt.Errorf("failed to find command: %w", err)
			}

			if savedCommand == nil {
				fmt.Printf("%s Command \"%s\" not found!\n", utils.Red("✗"), name)
				fmt.Printf("%s\n", utils.Blue("Use \"pock list\" to see all saved commands."))
				return nil
			}

			// Delete the command
			deleted, err := storage.DeleteSavedCommand(savedCommand.ID)
			if err != nil {
				return fmt.Errorf("failed to delete command: %w", err)
			}

			if deleted {
				fmt.Printf("%s Command \"%s\" removed successfully!\n", utils.Green("✓"), utils.GreenBold(name))
			} else {
				fmt.Printf("%s Failed to remove command \"%s\"\n", utils.Red("✗"), name)
			}

			return nil
		},
	}

	return cmd
}
