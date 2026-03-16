package cmd

import (
	"fmt"
	"os"
	"pock/internal/helpers"
	"pock/internal/storage"
	"pock/internal/utils"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

// lookupResult holds the result of a command database lookup.
type lookupResult struct {
	cmd *storage.SavedCommand
	err error
}

// lookupMsg is the Bubble Tea message emitted when the lookup finishes.
type lookupMsg lookupResult

// lookupModel is a minimal Bubble Tea model that animates a spinner while
// looking up a saved command by name in the background.
type lookupModel struct {
	spinner spinner.Model
	name    string
	result  *lookupResult
}

func (m lookupModel) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		func() tea.Msg {
			cmd, err := storage.GetSavedCommandByName(m.name)
			return lookupMsg{cmd: cmd, err: err}
		},
	)
}

func (m lookupModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case lookupMsg:
		r := lookupResult(msg)
		m.result = &r
		return m, tea.Quit
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m lookupModel) View() string {
	if m.result != nil {
		return ""
	}
	return fmt.Sprintf("%s Looking for command %q...\n", m.spinner.View(), m.name)
}

// findCommand runs a Bubble Tea spinner program while looking up a saved command.
func findCommand(name string) (*storage.SavedCommand, error) {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("6"))

	p := tea.NewProgram(lookupModel{spinner: s, name: name})
	finalModel, err := p.Run()
	if err != nil {
		// Fall back to a direct lookup when the TUI cannot run (e.g. non-TTY).
		return storage.GetSavedCommandByName(name)
	}

	if lm, ok := finalModel.(lookupModel); ok && lm.result != nil {
		return lm.result.cmd, lm.result.err
	}
	return storage.GetSavedCommandByName(name)
}

// NewRunCommand creates the run command
func NewRunCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "run <name>",
		Short:             "Run a saved command by name or execute a bash script file",
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: completeSavedCommandNames,
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]

			// Find the command with an animated spinner.
			savedCommand, err := findCommand(name)
			if err != nil {
				return fmt.Errorf("failed to find command: %w", err)
			}

			if savedCommand == nil {
				helpers.PrintCommandNotFound(name)
				return nil
			}

			// Check if command is trusted (for imported/marketplace commands)
			if !savedCommand.Trusted {
				fmt.Printf("%s %s\n", utils.Yellow("⚠"), utils.Yellow("Warning: This command is from an external source"))
				fmt.Printf("%s Source: %s\n", utils.Gray("ℹ"), utils.Cyan(savedCommand.Source))
				fmt.Printf("%s Command: %s\n", utils.Gray("ℹ"), savedCommand.Command)
				fmt.Printf("\n")

				// Ask for trust confirmation unless --yes flag is provided
				skipConfirm, _ := cmd.Flags().GetBool("yes")
				if !skipConfirm {
					fmt.Printf("%s Do you want to trust and run this command? [y/N]: ", utils.Yellow("?"))
					var response string
					fmt.Scanln(&response)
					response = strings.ToLower(strings.TrimSpace(response))

					if response != "y" && response != "yes" {
						fmt.Printf("%s Command execution cancelled\n", utils.Red("✗"))
						return nil
					}

					// Mark as trusted for future runs
					if err := storage.MarkCommandAsTrusted(savedCommand.ID); err != nil {
						fmt.Printf("%s Warning: failed to mark command as trusted: %v\n", utils.Yellow("⚠"), err)
					} else {
						fmt.Printf("%s Command marked as trusted\n", utils.Green("✓"))
					}
				}
				fmt.Println()
			}

			fmt.Printf("%s Found command: %s\n", utils.Green("✓"), utils.Yellow(savedCommand.Command))
			fmt.Printf("%s\n", utils.Gray("─────────────────────────────────────────────────"))

			// Execute the command
			result := utils.ExecuteCommandInteractive(savedCommand.Command)

			fmt.Printf("%s\n", utils.Gray("─────────────────────────────────────────────────"))

			// Determine if we should log output (privacy control)
			noLogOutput, _ := cmd.Flags().GetBool("no-log-output")
			// Check environment variable if flag is not set
			if !noLogOutput {
				envValue := strings.ToLower(os.Getenv("POCK_HISTORY_LOG"))
				if envValue == "0" || envValue == "false" || envValue == "no" {
					noLogOutput = true
				}
			}

			// Save to history
			status := "success"
			logOutput := ""
			if !noLogOutput {
				logOutput = result.Output
				if !result.Success && result.Error != "" {
					logOutput = result.Error
				}
			}
			if !result.Success {
				status = "failure"
			}

			_, err = storage.CreateCommandHistory(
				savedCommand.Name,
				savedCommand.Command,
				status,
				logOutput,
				result.ExitCode,
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

	cmd.Flags().Bool("no-log-output", false, "Don't save command output/error to history (privacy)")
	cmd.Flags().BoolP("yes", "y", false, "Skip trust confirmation for untrusted commands")

	return cmd
}
