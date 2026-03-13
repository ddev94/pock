package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"pock/internal/storage"
	"pock/internal/utils"
	"strings"

	"github.com/spf13/cobra"
)

// NewImportCommand creates the import command
func NewImportCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "import <file-or-url>",
		Short: "Import commands from a JSON file or URL",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			source := args[0]
			force, _ := cmd.Flags().GetBool("force")

			// Read the data
			var data []byte
			var err error

			if strings.HasPrefix(source, "http://") || strings.HasPrefix(source, "https://") {
				// Fetch from URL
				fmt.Printf("%s Fetching from %s...\n", utils.Cyan("→"), source)
				resp, err := http.Get(source)
				if err != nil {
					return fmt.Errorf("failed to fetch URL: %w", err)
				}
				defer resp.Body.Close()

				if resp.StatusCode != http.StatusOK {
					return fmt.Errorf("failed to fetch URL: status code %d", resp.StatusCode)
				}

				data, err = io.ReadAll(resp.Body)
				if err != nil {
					return fmt.Errorf("failed to read response: %w", err)
				}
			} else {
				// Read from file
				data, err = os.ReadFile(source)
				if err != nil {
					return fmt.Errorf("failed to read file: %w", err)
				}
			}

			// Parse the commands
			var imported []storage.ExportedCommand
			err = json.Unmarshal(data, &imported)
			if err != nil {
				return fmt.Errorf("failed to parse commands: %w", err)
			}

			if len(imported) == 0 {
				fmt.Printf("%s\n", utils.Yellow("No commands to import."))
				return nil
			}

			fmt.Printf("%s Found %d command(s) to import\n", utils.Cyan("→"), len(imported))

			// Import commands
			var importedCount int
			var skippedCount int

			for _, cmd := range imported {
				// Check if command already exists
				existing, err := storage.GetSavedCommandByName(cmd.Name)
				if err != nil {
					return fmt.Errorf("failed to check existing commands: %w", err)
				}

				if existing != nil && !force {
					fmt.Printf("%s Skipping \"%s\" (already exists)\n", utils.Yellow("⊘"), cmd.Name)
					skippedCount++
					continue
				}

				if existing != nil && force {
					// Delete existing command
					_, err := storage.DeleteSavedCommand(existing.ID)
					if err != nil {
						return fmt.Errorf("failed to delete existing command: %w", err)
					}
				}

				// Create the command
				_, err = storage.CreateSavedCommand(storage.NewSavedCommandInput{
					Name:        cmd.Name,
					Command:     cmd.Command,
					Description: cmd.Description,
				})

				if err != nil {
					return fmt.Errorf("failed to import command: %w", err)
				}

				fmt.Printf("%s Imported \"%s\"\n", utils.Green("✓"), cmd.Name)
				importedCount++
			}

			fmt.Printf("\n%s Successfully imported %d command(s)",
				utils.Green("✓"),
				importedCount)
			if skippedCount > 0 {
				fmt.Printf(" (skipped %d)", skippedCount)
			}
			fmt.Println()

			return nil
		},
	}

	cmd.Flags().BoolP("force", "f", false, "Overwrite existing commands")

	return cmd
}
