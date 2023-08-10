package list

import (
	"encoding/json"
	"fmt"

	"github.com/moutend/clickupctl/internal/constant"
	"github.com/raksul/go-clickup/clickup"
	"github.com/spf13/cobra"
)

var viewCommand = &cobra.Command{
	Use:     "view",
	Aliases: []string{"v"},
	RunE:    viewCommandRunE,
}

func viewCommandRunE(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return nil
	}

	client, ok := cmd.Context().Value(constant.ClientContextKey).(*clickup.Client)

	if !ok {
		return fmt.Errorf("failed to get client")
	}

	listID := args[0]

	list, _, err := client.Lists.GetList(cmd.Context(), listID)

	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(list, "", "  ")

	if err != nil {
		return err
	}

	cmd.Printf("%s\n", data)

	return nil
}

func init() {
	Command.AddCommand(viewCommand)
}
