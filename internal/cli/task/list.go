package task

import (
	"encoding/json"
	"fmt"

	"github.com/moutend/clickupctl/internal/constant"
	"github.com/raksul/go-clickup/clickup"
	"github.com/spf13/cobra"
)

var listCommand = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	RunE:    listCommandRunE,
}

func listCommandRunE(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return nil
	}

	listID := args[0]

	archived, _ := cmd.Flags().GetBool("archived")
	page, _ := cmd.Flags().GetInt("page")
	orderby, _ := cmd.Flags().GetString("orderby")
	reverse, _ := cmd.Flags().GetBool("reverse")
	subtasks, _ := cmd.Flags().GetBool("subtasks")
	statuses, _ := cmd.Flags().GetStringSlice("status")
	includeClosed, _ := cmd.Flags().GetBool("include-closed")
	assignees, _ := cmd.Flags().GetStringSlice("assignee")

	option := &clickup.GetTasksOptions{
		Archived:      archived,
		Page:          page,
		OrderBy:       orderby,
		Reverse:       reverse,
		Subtasks:      subtasks,
		Statuses:      statuses,
		IncludeClosed: includeClosed,
		Assignees:     assignees,
	}

	client, ok := cmd.Context().Value(constant.ClientContextKey).(*clickup.Client)

	if !ok {
		return fmt.Errorf("failed to get client")
	}

	tasks, _, err := client.Tasks.GetTasks(cmd.Context(), listID, option)

	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(tasks, "", "  ")

	if err != nil {
		return err
	}

	cmd.Printf("%s\n", data)

	return nil
}

func init() {
	Command.AddCommand(listCommand)

	listCommand.PersistentFlags().BoolP("archived", "a", false, "archived")
	listCommand.PersistentFlags().IntP("page", "p", 0, "page")
	listCommand.PersistentFlags().StringP("orderby", "o", "", "orderby")
	listCommand.PersistentFlags().BoolP("reverse", "r", false, "reverse")
	listCommand.PersistentFlags().BoolP("subtasks", "", false, "subtask")
	listCommand.PersistentFlags().StringSliceP("status", "s", nil, "status")
	listCommand.PersistentFlags().BoolP("include-closed", "", false, "closed")
	listCommand.PersistentFlags().StringSliceP("assignee", "", nil, "assignees")
	listCommand.PersistentFlags().StringP("due-after", "", "", "due-after")
	listCommand.PersistentFlags().StringP("due-before", "", "", "due-before")
	listCommand.PersistentFlags().StringP("create-after", "", "", "create-after")
	listCommand.PersistentFlags().StringP("create-before", "", "", "create-before")
	listCommand.PersistentFlags().StringP("update-after", "", "", "update-after")
	listCommand.PersistentFlags().StringP("update-before", "", "", "update-before")
}
