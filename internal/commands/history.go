package commands

import (
	"fmt"
	"pock/internal/storage"
	"pock/internal/utils"
	"strings"

	"github.com/spf13/cobra"
)

// NewHistoryCommand creates the history command
func NewHistoryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "history [command-name]",
		Short: "View command execution history",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			limit, _ := cmd.Flags().GetInt("limit")
			clear, _ := cmd.Flags().GetBool("clear")
			showOutput, _ := cmd.Flags().GetBool("output")

			if clear {
				if len(args) > 0 {
					// Clear history for specific command
					commandName := args[0]
					err := storage.ClearCommandHistoryByName(commandName)
					if err != nil {
						return fmt.Errorf("failed to clear history: %w", err)
					}
					fmt.Printf("%s History for command \"%s\" cleared successfully!\n", utils.Green("✓"), commandName)
				} else {
					// Clear all history
					err := storage.ClearCommandHistory()
					if err != nil {
						return fmt.Errorf("failed to clear history: %w", err)
					}
					fmt.Printf("%s Command history cleared successfully!\n", utils.Green("✓"))
				}
				return nil
			}

			var histories []storage.CommandHistory
			var err error
			var title string

			// If command name is provided, filter by that command
			if len(args) > 0 {
				commandName := args[0]
				histories, err = storage.GetCommandHistoryByName(commandName, limit)
				if err != nil {
					return fmt.Errorf("failed to get history: %w", err)
				}
				title = fmt.Sprintf("Command History for \"%s\" (%d):", commandName, len(histories))
			} else {
				histories, err = storage.GetCommandHistory(limit)
				if err != nil {
					return fmt.Errorf("failed to get history: %w", err)
				}
				title = fmt.Sprintf("Command History (%d):", len(histories))
			}

			if len(histories) == 0 {
				if len(args) > 0 {
					fmt.Printf("%s\n", utils.Yellow(fmt.Sprintf("No history found for command \"%s\".", args[0])))
				} else {
					fmt.Printf("%s\n", utils.Yellow("No command history found."))
				}
				return nil
			}

			fmt.Printf("\n%s\n\n", utils.CyanBold(title))

			if showOutput {
				// Show detailed view with output
				for i, h := range histories {
					statusColor := utils.Green
					statusSymbol := "✓"
					if h.Status == "failure" {
						statusColor = utils.Red
						statusSymbol = "✗"
					}

					fmt.Printf("%s %s\n", statusColor(statusSymbol), utils.CyanBold(h.CommandName))
					fmt.Printf("  %s %s\n", utils.Gray("Date:"), h.Date.Format("2006-01-02 15:04:05"))
					fmt.Printf("  %s %s\n", utils.Gray("Command:"), utils.Yellow(h.CommandText))
					fmt.Printf("  %s %s\n", utils.Gray("Status:"), statusColor(h.Status))
					fmt.Printf("  %s %s\n", utils.Gray("Duration:"), utils.FormatDuration(h.ExecutionTime))

					if h.Log != "" {
						fmt.Printf("  %s\n", utils.Gray("Output:"))
						// Indent the output
						lines := strings.Split(h.Log, "\n")
						for _, line := range lines {
							if line != "" {
								fmt.Printf("    %s\n", line)
							}
						}
					}

					if i < len(histories)-1 {
						fmt.Println()
					}
				}
			} else {
				// Show table view
				headers := []string{"Date", "Command", "Status", "Time"}
				var rows [][]string

				for _, h := range histories {
					statusColor := utils.Green
					if h.Status == "failure" {
						statusColor = utils.Red
					}

					row := []string{
						h.Date.Format("2006-01-02 15:04:05"),
						utils.Yellow(h.CommandName),
						statusColor(h.Status),
						utils.FormatDuration(h.ExecutionTime),
					}
					rows = append(rows, row)
				}

				utils.RenderTable(headers, rows)
			}

			return nil
		},
	}

	cmd.Flags().IntP("limit", "l", 20, "Limit the number of history entries")
	cmd.Flags().Bool("clear", false, "Clear all command history")
	cmd.Flags().BoolP("output", "o", false, "Show command output in history")

	return cmd
}
