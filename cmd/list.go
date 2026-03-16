package cmd

import (
	"fmt"
	"pock/internal/storage"
	"pock/internal/utils"

	"github.com/spf13/cobra"
)

// NewListCommand creates the list command
func NewListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List all saved commands",
		RunE: func(cmd *cobra.Command, args []string) error {
			showStats, _ := cmd.Flags().GetBool("stats")

			commands, err := storage.GetAllSavedCommands()
			if err != nil {
				return fmt.Errorf("failed to get commands: %w", err)
			}

			if len(commands) == 0 {
				fmt.Printf("%s\n", utils.Yellow("No saved commands found."))
				fmt.Printf("%s\n", utils.Blue("Add commands with: pock add <name> \"<command>\""))
				return nil
			}

			// Get settings
			settings, err := storage.GetSettings()
			if err != nil {
				return fmt.Errorf("failed to get settings: %w", err)
			}

			if settings.ListLayout == "simple" {
				// Simple layout
				fmt.Printf("\n%s\n\n", utils.CyanBold(fmt.Sprintf("Saved Commands (%d):", len(commands))))

				for _, cmd := range commands {
					fmt.Printf("%s %s\n", utils.Green("•"), utils.GreenBold(cmd.Name))
					fmt.Printf("  %s %s\n", utils.Gray("Command:"), utils.YellowBold(cmd.Command))
					if cmd.Description != "" {
						fmt.Printf("  %s %s\n", utils.Gray("Description:"), cmd.Description)
					}
					if showStats {
						stats, err := storage.GetCommandStats(cmd.Name)
						if err == nil && stats.TotalRuns > 0 {
							fmt.Printf("  %s Total: %d | Success: %d | Failed: %d\n",
								utils.Gray("Stats:"),
								stats.TotalRuns,
								stats.SuccessfulRuns,
								stats.FailedRuns)
						}
					}
					fmt.Println()
				}
			} else {
				// Table layout
				fmt.Printf("\n%s\n\n", utils.CyanBold(fmt.Sprintf("Saved Commands (%d):", len(commands))))

				headers := []string{"Name", "Command", "Description"}
				if showStats {
					headers = append(headers, "Runs", "Success", "Failed")
				}

				var rows [][]string
				for _, cmd := range commands {
					row := []string{
						utils.Green(cmd.Name),
						utils.Yellow(cmd.Command),
						cmd.Description,
					}

					if showStats {
						stats, err := storage.GetCommandStats(cmd.Name)
						if err == nil {
							row = append(row,
								fmt.Sprintf("%d", stats.TotalRuns),
								fmt.Sprintf("%d", stats.SuccessfulRuns),
								fmt.Sprintf("%d", stats.FailedRuns))
						} else {
							row = append(row, "0", "0", "0")
						}
					}

					rows = append(rows, row)
				}

				utils.RenderTable(headers, rows)
			}

			return nil
		},
	}

	cmd.Flags().BoolP("stats", "s", false, "Show execution statistics for each command")

	return cmd
}
