package task

import (
	"encoding/json"
	"fmt"

	"github.com/moutend/clickupctl/internal/constant"
	"github.com/raksul/go-clickup/clickup"
	"github.com/spf13/cobra"
)

var createCommand = &cobra.Command{
	Use:     "create",
	Aliases: []string{"c"},
	RunE:    createCommandRunE,
}

func createCommandRunE(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return nil
	}

	client, ok := cmd.Context().Value(constant.ClientContextKey).(*clickup.Client)

	if !ok {
		return fmt.Errorf("failed to get client")
	}

	listID := args[0]
	taskRequest := buildTaskRequest(cmd)

	task, _, err := client.Tasks.CreateTask(cmd.Context(), listID, taskRequest)

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
	Command.AddCommand(createCommand)

	setupTaskRequestFlags(createCommand)
}
