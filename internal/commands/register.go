package commands

import "github.com/spf13/cobra"

func registerCommands(rootCmd *cobra.Command) {
	rootCmd.AddCommand(
		NewAddCommand(),
		NewListCommand(),
		NewRunCommand(),
		NewRemoveCommand(),
		NewHistoryCommand(),
		NewExportCommand(),
		NewImportCommand(),
		NewConfigCommand(),
		NewInstallCommand(),
		NewBrowseCommand(),
		NewPublishCommand(),
	)
}
