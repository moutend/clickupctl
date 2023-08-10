package task

import (
	"encoding/json"
	"fmt"

	"github.com/moutend/clickupctl/internal/constant"
	"github.com/raksul/go-clickup/clickup"
	"github.com/spf13/cobra"
)

var updateCommand = &cobra.Command{
	Use:     "update",
	Aliases: []string{"u"},
	RunE:    updateCommandRunE,
}

func updateCommandRunE(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return nil
	}

	client, ok := cmd.Context().Value(constant.ClientContextKey).(*clickup.Client)

	if !ok {
		return fmt.Errorf("failed to get client")
	}

	listID := args[0]
	getTaskOptions := buildGetTaskOptions(cmd)
	taskRequest := buildTaskRequest(cmd)

	task, _, err := client.Tasks.UpdateTask(cmd.Context(), listID, getTaskOptions, taskRequest)

	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(task, "", "  ")

	if err != nil {
		return err
	}

	cmd.Printf("%s\n", data)

	return nil
}

func init() {
	Command.AddCommand(updateCommand)

	setupGetTaskOptionsFlags(updateCommand)
	setupTaskRequestFlags(updateCommand)
}
