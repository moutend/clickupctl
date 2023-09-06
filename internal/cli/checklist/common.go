package checklist

import (
	"github.com/raksul/go-clickup/clickup"
	"github.com/spf13/cobra"
)

func setupChecklistItemRequestFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP("name", "", "", "name")
	cmd.PersistentFlags().IntP("assignee", "", 0, "assignee")
	cmd.PersistentFlags().BoolP("resolved", "", false, "resolved")
}

func buildChecklistItemRequest(cmd *cobra.Command) *clickup.ChecklistItemRequest {
	var (
		name     string
		assignee int
		resolved bool
	)

	if value, _ := cmd.Flags().GetString("name"); value != "" {
		name = value
	}
	if value, _ := cmd.Flags().GetInt("assignee"); value != 0 {
		assignee = value
	}
	if value, _ := cmd.Flags().GetBool("resolved"); value {
		resolved = value
	}

	return &clickup.ChecklistItemRequest{
		Name:     name,
		Assignee: assignee,
		Resolved: resolved,
	}
}
