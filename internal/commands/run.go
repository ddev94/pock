package commands

import (
	"fmt"
	"pock/internal/helpers"
	"pock/internal/storage"
	"pock/internal/utils"

	"github.com/spf13/cobra"
)

// NewRunCommand creates the run command
func NewRunCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "run <name>",
		Short:             "Run a saved command by name or execute a bash script file",
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: completeSavedCommandNames,
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]

			fmt.Printf("%s Looking for command \"%s\"...\n", utils.Cyan("→"), name)

			// Find the command
			savedCommand, err := storage.GetSavedCommandByName(name)
			if err != nil {
				return fmt.Errorf("failed to find command: %w", err)
			}

			if savedCommand == nil {
				helpers.PrintCommandNotFound(name)
				return nil
			}

			fmt.Printf("%s Found command: %s\n", utils.Green("✓"), utils.Yellow(savedCommand.Command))
			fmt.Printf("%s\n", utils.Gray("─────────────────────────────────────────────────"))

			// Execute the command
			result := utils.ExecuteCommandInteractive(savedCommand.Command)

			fmt.Printf("%s\n", utils.Gray("─────────────────────────────────────────────────"))

			// Save to history
			status := "success"
			logOutput := result.Output
			if !result.Success {
				status = "failure"
				if result.Error != "" {
					logOutput = result.Error
				}
			}

			_, err = storage.CreateCommandHistory(
				savedCommand.Name,
				savedCommand.Command,
				status,
				logOutput,
				result.ExecutionTime,
			)
			if err != nil {
				fmt.Printf("%s Failed to save command history: %v\n", utils.Yellow("⚠"), err)
			}

			// Print execution info
			if result.Success {
				fmt.Printf("%s Command executed successfully in %s\n",
					utils.Green("✓"),
					utils.Cyan(utils.FormatDuration(result.ExecutionTime)))
			} else {
				fmt.Printf("%s Command failed with exit code %d in %s\n",
					utils.Red("✗"),
					result.ExitCode,
					utils.Cyan(utils.FormatDuration(result.ExecutionTime)))
			}

			if !result.Success {
				return fmt.Errorf("command execution failed")
			}

			return nil
		},
	}

	return cmd
}
