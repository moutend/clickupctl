package task

import (
	"github.com/raksul/go-clickup/clickup"
	"github.com/spf13/cobra"
)

func setupGetTaskOptionsFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP("custom-task-ids", "", "", "custom-task-ids")
	cmd.PersistentFlags().IntP("team-id", "", 0, "team-id")
	cmd.PersistentFlags().BoolP("include-subtasks", "", false, "include-subtasks")
}

func setupTaskRequestFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP("name", "", "", "name")
	cmd.PersistentFlags().StringP("description", "", "", "description")
	cmd.PersistentFlags().IntSliceP("assignees", "", nil, "assignees")
	cmd.PersistentFlags().StringSliceP("tags", "", nil, "tags")
	cmd.PersistentFlags().StringP("status", "", "", "status")
	cmd.PersistentFlags().IntP("priority", "", 0, "priority")
	cmd.PersistentFlags().IntP("time-estimate", "", 0, "time-estimate")
	cmd.PersistentFlags().BoolP("notify-all", "", false, "notify-all")
	cmd.PersistentFlags().StringP("parent", "", "", "parent")
	cmd.PersistentFlags().StringP("links-to", "", "", "links-to")
	cmd.PersistentFlags().BoolP("check-required-custom-fields", "", false, "check-required-custom-fields")
}

func buildGetTaskOptions(cmd *cobra.Command) *clickup.GetTaskOptions {
	var (
		customTaskIDs   string
		teamID          int
		includeSubTasks bool
	)

	if value, _ := cmd.Flags().GetString("custom-task-ids"); value != "" {
		customTaskIDs = value
	}
	if value, _ := cmd.Flags().GetInt("team-id"); value != 0 {
		teamID = value
	}
	if value, _ := cmd.Flags().GetBool("include-sub-tasks"); value {
		includeSubTasks = value
	}

	return &clickup.GetTaskOptions{
		CustomTaskIDs:   customTaskIDs,
		TeamID:          teamID,
		IncludeSubTasks: includeSubTasks,
	}
}

func buildTaskRequest(cmd *cobra.Command) *clickup.TaskRequest {
	var (
		name                      string
		description               string
		assignees                 []int
		tags                      []string
		status                    string
		priority                  int
		dueDate                   *clickup.Date
		dueDateTime               bool
		timeEstimate              int
		startDate                 *clickup.Date
		startDateTime             bool
		notifyAll                 bool
		parent                    string
		linksTo                   string
		checkRequiredCustomFields bool
		customFields              []clickup.CustomFieldInTaskRequest
	)

	if value, _ := cmd.Flags().GetString("name"); value != "" {
		name = value
	}
	if value, _ := cmd.Flags().GetString("description"); value != "" {
		description = value
	}
	if value, _ := cmd.Flags().GetIntSlice("assignees"); value != nil {
		assignees = value
	}
	if value, _ := cmd.Flags().GetStringSlice("tags"); value != nil {
		tags = value
	}
	if value, _ := cmd.Flags().GetString("status"); value != "" {
		status = value
	}
	if value, _ := cmd.Flags().GetInt("priority"); value != 0 {
		priority = value
	}
	if value, _ := cmd.Flags().GetInt("time-estimate"); value != 0 {
		timeEstimate = value
	}
	if value, _ := cmd.Flags().GetBool("notify-all"); value {
		notifyAll = value
	}
	if value, _ := cmd.Flags().GetString("parent"); value != "" {
		parent = value
	}
	if value, _ := cmd.Flags().GetString("links-to"); value != "" {
		linksTo = value
	}
	if value, _ := cmd.Flags().GetBool("check-required-custom-fields"); value {
		checkRequiredCustomFields = value
	}

	return &clickup.TaskRequest{
		Name:                      name,
		Description:               description,
		Assignees:                 assignees,
		Tags:                      tags,
		Status:                    status,
		Priority:                  priority,
		DueDate:                   dueDate,
		DueDateTime:               dueDateTime,
		TimeEstimate:              timeEstimate,
		StartDate:                 startDate,
		StartDateTime:             startDateTime,
		NotifyAll:                 notifyAll,
		Parent:                    parent,
		LinksTo:                   linksTo,
		CheckRequiredCustomFields: checkRequiredCustomFields,
		CustomFields:              customFields,
	}
}
