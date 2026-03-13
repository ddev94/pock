package commands

import (
	"pock/internal/storage"
	"strings"

	"github.com/spf13/cobra"
)

// completeSavedCommandNames provides shell completion for saved command names.
func completeSavedCommandNames(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if len(args) > 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	savedCommands, err := storage.GetAllSavedCommands()
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	needle := strings.ToLower(toComplete)
	results := make([]string, 0, len(savedCommands))
	for _, saved := range savedCommands {
		if needle == "" || strings.HasPrefix(strings.ToLower(saved.Name), needle) {
			results = append(results, saved.Name)
		}
	}

	return results, cobra.ShellCompDirectiveNoFileComp
}
