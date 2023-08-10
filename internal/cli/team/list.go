package team

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
	client, ok := cmd.Context().Value(constant.ClientContextKey).(*clickup.Client)

	if !ok {
		return fmt.Errorf("failed to get client")
	}

	teams, _, err := client.Teams.GetTeams(cmd.Context())

	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(teams, "", "  ")

	if err != nil {
		return err
	}

	cmd.Printf("%s\n", data)

	return nil
}

func init() {
	Command.AddCommand(listCommand)
}
