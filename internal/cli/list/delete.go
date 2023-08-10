package list

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

	listID := args[0]

	_, err := client.Lists.DeleteList(cmd.Context(), listID)

	if err != nil {
		return err
	}

	cmd.Printf("%s\n", listID)

	return nil
}

func init() {
	Command.AddCommand(deleteCommand)
}
