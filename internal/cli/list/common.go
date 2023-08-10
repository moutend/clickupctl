package list

import (
	"github.com/raksul/go-clickup/clickup"
	"github.com/spf13/cobra"
)

func setupListRequestFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP("name", "", "", "name")
	cmd.PersistentFlags().StringP("content", "", "", "content")
	cmd.PersistentFlags().IntP("priority", "", 0, "priority")
	cmd.PersistentFlags().IntP("assignee", "", 0, "assignee")
	cmd.PersistentFlags().StringP("status", "", "", "status")
}

func buildListRequest(cmd *cobra.Command) *clickup.ListRequest {
	var (
		name        string
		content     string
		dueDate     *clickup.Date
		dueDateTime bool
		priority    int
		assignee    int
		status      string
	)

	if value, _ := cmd.Flags().GetString("name"); value != "" {
		name = value
	}
	if value, _ := cmd.Flags().GetString("content"); value != "" {
		content = value
	}
	if value, _ := cmd.Flags().GetInt("priority"); value != 0 {
		priority = value
	}
	if value, _ := cmd.Flags().GetInt("assignee"); value != 0 {
		assignee = value
	}
	if value, _ := cmd.Flags().GetString("status"); value != "" {
		status = value
	}

	return &clickup.ListRequest{
		Name:        name,
		Content:     content,
		DueDate:     dueDate,
		DueDateTime: dueDateTime,
		Priority:    priority,
		Assignee:    assignee,
		Status:      status,
	}
}
