package cmd

import "github.com/spf13/cobra"

func registerCommands(rootCmd *cobra.Command) {
	if EnableAddCommand {
		rootCmd.AddCommand(NewAddCommand())
	}
	if EnableListCommand {
		rootCmd.AddCommand(NewListCommand())
	}
	if EnableRunCommand {
		rootCmd.AddCommand(NewRunCommand())
	}
	if EnableRemoveCommand {
		rootCmd.AddCommand(NewRemoveCommand())
	}
	if EnableHistoryCommand {
		rootCmd.AddCommand(NewHistoryCommand())
	}
	if EnableExportCommand {
		rootCmd.AddCommand(NewExportCommand())
	}
	if EnableImportCommand {
		rootCmd.AddCommand(NewImportCommand())
	}
	if EnableInstallCommand {
		rootCmd.AddCommand(NewInstallCommand())
	}
	if EnableBrowseCommand {
		rootCmd.AddCommand(NewBrowseCommand())
	}
	if EnablePublishCommand {
		rootCmd.AddCommand(NewPublishCommand())
	}
}
