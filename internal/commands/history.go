package commands

import (
	"fmt"
	"pock/internal/storage"
	"pock/internal/utils"

	"github.com/spf13/cobra"
)

// NewHistoryCommand creates the history command
func NewHistoryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "history",
		Short: "View command execution history",
		RunE: func(cmd *cobra.Command, args []string) error {
			limit, _ := cmd.Flags().GetInt("limit")
			clear, _ := cmd.Flags().GetBool("clear")

			if clear {
				err := storage.ClearCommandHistory()
				if err != nil {
					return fmt.Errorf("failed to clear history: %w", err)
				}
				fmt.Printf("%s Command history cleared successfully!\n", utils.Green("✓"))
				return nil
			}

			histories, err := storage.GetCommandHistory(limit)
			if err != nil {
				return fmt.Errorf("failed to get history: %w", err)
			}

			if len(histories) == 0 {
				fmt.Printf("%s\n", utils.Yellow("No command history found."))
				return nil
			}

			fmt.Printf("\n%s\n\n", utils.CyanBold(fmt.Sprintf("Command History (%d):", len(histories))))

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

			return nil
		},
	}

	cmd.Flags().IntP("limit", "l", 20, "Limit the number of history entries")
	cmd.Flags().Bool("clear", false, "Clear all command history")

	return cmd
}
