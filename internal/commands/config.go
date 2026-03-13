package commands

import (
	"fmt"
	"pock/internal/storage"
	"pock/internal/utils"

	"github.com/spf13/cobra"
)

// NewConfigCommand creates the config command
func NewConfigCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manage configuration settings",
	}

	// config list
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List all configuration settings",
		RunE: func(cmd *cobra.Command, args []string) error {
			settings, err := storage.GetSettings()
			if err != nil {
				return fmt.Errorf("failed to get settings: %w", err)
			}

			fmt.Printf("\n%s\n\n", utils.CyanBold("Configuration Settings:"))
			fmt.Printf("%s %s\n", utils.Green("listLayout:"), settings.ListLayout)
			fmt.Printf("%s %s\n", utils.Green("dateFormat:"), settings.DateFormat)

			return nil
		},
	}

	// config get
	getCmd := &cobra.Command{
		Use:   "get <key>",
		Short: "Get a configuration value",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			key := args[0]
			settings, err := storage.GetSettings()
			if err != nil {
				return fmt.Errorf("failed to get settings: %w", err)
			}

			var value string
			switch key {
			case "listLayout":
				value = settings.ListLayout
			case "dateFormat":
				value = settings.DateFormat
			default:
				fmt.Printf("%s Unknown setting key: %s\n", utils.Red("✗"), key)
				fmt.Printf("%s Valid keys: listLayout, dateFormat\n", utils.Blue("ℹ"))
				return nil
			}

			fmt.Printf("%s\n", utils.Green(value))
			return nil
		},
	}

	// config set
	setCmd := &cobra.Command{
		Use:   "set <key> <value>",
		Short: "Set a configuration value",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			key := args[0]
			value := args[1]

			// Validate the setting
			if key == "listLayout" && value != "table" && value != "simple" {
				fmt.Printf("%s Invalid value for listLayout. Must be 'table' or 'simple'\n", utils.Red("✗"))
				return nil
			}
			if key == "dateFormat" && value != "relative" && value != "locale" && value != "iso" {
				fmt.Printf("%s Invalid value for dateFormat. Must be 'relative', 'locale', or 'iso'\n", utils.Red("✗"))
				return nil
			}

			updates := map[string]string{key: value}
			_, err := storage.UpdateSettings(updates)
			if err != nil {
				return fmt.Errorf("failed to update settings: %w", err)
			}

			fmt.Printf("%s Configuration updated: %s = %s\n", utils.Green("✓"), utils.Cyan(key), utils.Yellow(value))
			return nil
		},
	}

	// config reset
	resetCmd := &cobra.Command{
		Use:   "reset",
		Short: "Reset configuration to defaults",
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := storage.ResetSettings()
			if err != nil {
				return fmt.Errorf("failed to reset settings: %w", err)
			}

			fmt.Printf("%s Configuration reset to defaults!\n", utils.Green("✓"))
			return nil
		},
	}

	cmd.AddCommand(listCmd, getCmd, setCmd, resetCmd)

	return cmd
}
