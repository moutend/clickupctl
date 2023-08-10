package list

import (
	"encoding/json"
	"fmt"

	"github.com/moutend/clickupctl/internal/constant"
	"github.com/raksul/go-clickup/clickup"
	"github.com/spf13/cobra"
)

var memberCommand = &cobra.Command{
	Use:     "member",
	Aliases: []string{"m"},
	RunE:    memberCommandRunE,
}

func memberCommandRunE(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return nil
	}

	client, ok := cmd.Context().Value(constant.ClientContextKey).(*clickup.Client)

	if !ok {
		return fmt.Errorf("failed to get client")
	}

	listID := args[0]

	members, _, err := client.Members.GetListMembers(cmd.Context(), listID)

	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(members, "", "  ")

	if err != nil {
		return err
	}

	cmd.Printf("%s\n", data)

	return nil
}

func init() {
	Command.AddCommand(memberCommand)
}
