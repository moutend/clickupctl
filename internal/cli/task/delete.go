package task

import (
	"fmt"

	"github.com/moutend/clickupctl/internal/constant"
	"github.com/raksul/go-clickup/clickup"
	"github.com/spf13/cobra"
)

var deleteCommand = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"d"},
	RunE:    deleteCommandRunE,
}

func deleteCommandRunE(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return nil
	}

	client, ok := cmd.Context().Value(constant.ClientContextKey).(*clickup.Client)

	if !ok {
		return fmt.Errorf("failed to get client")
	}

	taskID := args[0]

	customTaskIDs, _ := cmd.Flags().GetString("custom-task-ids")
	teamID, _ := cmd.Flags().GetInt("team-id")
	includeSubTasks, _ := cmd.Flags().GetBool("include-sub-tasks")

	option := &clickup.GetTaskOptions{
		CustomTaskIDs:   customTaskIDs,
		TeamID:          teamID,
		IncludeSubTasks: includeSubTasks,
	}

	_, err := client.Tasks.DeleteTask(cmd.Context(), taskID, option)

	if err != nil {
		return err
	}

	cmd.Printf("%s\n", taskID)

	return nil
}

func init() {
	Command.AddCommand(deleteCommand)

	deleteCommand.PersistentFlags().StringP("custom-task-ids", "", "", "custom-task-ids")
	deleteCommand.PersistentFlags().IntP("team-id", "", 0, "team-id")
	deleteCommand.PersistentFlags().BoolP("include-subtasks", "", false, "include-subtasks")
}
