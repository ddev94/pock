package commands

import (
	"fmt"
	"pock/internal/storage"
	"pock/internal/utils"

	"github.com/spf13/cobra"
)

// NewAddCommand creates the add command
func NewAddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add <name> <command>",
		Short: "Add a new command to the command manager",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			commandText := args[1]
			description, _ := cmd.Flags().GetString("description")

			// Check if command already exists
			existing, err := storage.GetSavedCommandByName(name)
			if err != nil {
				return fmt.Errorf("failed to check existing commands: %w", err)
			}

			if existing != nil {
				fmt.Printf("%s Command with name \"%s\" already exists!\n", utils.Red("✗"), name)
				fmt.Printf("%s %s\n", utils.Yellow("Existing command:"), existing.Command)
				fmt.Printf("%s\n", utils.Blue("Use a different name or delete the existing command first."))
				return nil
			}

			// Create the command
			savedCommand, err := storage.CreateSavedCommand(storage.NewSavedCommandInput{
				Name:        name,
				Command:     commandText,
				Description: description,
			})

			if err != nil {
				return fmt.Errorf("failed to add command: %w", err)
			}

			fmt.Printf("%s Command \"%s\" added successfully!\n", utils.Green("✓"), utils.GreenBold(savedCommand.Name))
			fmt.Printf("%s %s\n", utils.Cyan("Command:"), utils.YellowBold(savedCommand.Command))
			if savedCommand.Description != "" {
				fmt.Printf("%s %s\n", utils.Gray("Description:"), savedCommand.Description)
			}
			fmt.Printf("\n%s\n", utils.Blue(fmt.Sprintf("You can now use it with: pock run %s", savedCommand.Name)))

			return nil
		},
	}

	cmd.Flags().StringP("description", "d", "", "Optional description for the command")

	return cmd
}
