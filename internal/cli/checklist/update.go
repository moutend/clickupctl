package checklist

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
	if len(args) < 2 {
		return nil
	}

	client, ok := cmd.Context().Value(constant.ClientContextKey).(*clickup.Client)

	if !ok {
		return fmt.Errorf("failed to get client")
	}

	checklistID := args[0]
	checklistItemID := args[1]
	request := buildChecklistItemRequest(cmd)

	checklist, _, err := client.Checklists.EditChecklistItem(cmd.Context(), checklistID, checklistItemID, request)

	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(checklist, "", "  ")

	if err != nil {
		return err
	}

	cmd.Printf("%s\n", data)

	return nil
}

func init() {
	Command.AddCommand(updateCommand)

	setupChecklistItemRequestFlags(updateCommand)
}
