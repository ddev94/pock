package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"pock/internal/storage"
	"pock/internal/utils"

	"github.com/spf13/cobra"
)

// NewExportCommand creates the export command
func NewExportCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export <output-file>",
		Short: "Export commands to a JSON file",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			outputFile := args[0]
			commandName, _ := cmd.Flags().GetString("name")
			author, _ := cmd.Flags().GetString("author")
			tags, _ := cmd.Flags().GetStringSlice("tags")
			version, _ := cmd.Flags().GetString("version")

			var commands []storage.SavedCommand
			var err error

			if commandName != "" {
				// Export specific command
				savedCmd, err := storage.GetSavedCommandByName(commandName)
				if err != nil {
					return fmt.Errorf("failed to find command: %w", err)
				}
				if savedCmd == nil {
					fmt.Printf("%s Command \"%s\" not found!\n", utils.Red("✗"), commandName)
					return nil
				}
				commands = []storage.SavedCommand{*savedCmd}
			} else {
				// Export all commands
				commands, err = storage.GetAllSavedCommands()
				if err != nil {
					return fmt.Errorf("failed to get commands: %w", err)
				}
			}

			if len(commands) == 0 {
				fmt.Printf("%s\n", utils.Yellow("No commands to export."))
				return nil
			}

			// Convert to exported format
			var exported []storage.ExportedCommand
			for _, cmd := range commands {
				exported = append(exported, storage.ExportedCommand{
					Name:        cmd.Name,
					Command:     cmd.Command,
					Description: cmd.Description,
					Author:      author,
					Tags:        tags,
					Version:     version,
				})
			}

			// Write to file
			data, err := json.MarshalIndent(exported, "", "  ")
			if err != nil {
				return fmt.Errorf("failed to marshal commands: %w", err)
			}

			err = os.WriteFile(outputFile, data, 0644)
			if err != nil {
				return fmt.Errorf("failed to write file: %w", err)
			}

			fmt.Printf("%s Exported %d command(s) to %s\n",
				utils.Green("✓"),
				len(commands),
				utils.Cyan(outputFile))

			return nil
		},
	}

	cmd.Flags().StringP("name", "n", "", "Export specific command by name")
	cmd.Flags().StringP("author", "a", "", "Author name for exported commands")
	cmd.Flags().StringSliceP("tags", "t", []string{}, "Tags for exported commands")
	cmd.Flags().StringP("version", "v", "", "Version for exported commands")

	return cmd
}
