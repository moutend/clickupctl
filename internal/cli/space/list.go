package space

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

	teamID := args[0]
	archived, _ := cmd.Flags().GetBool("archived")

	client, ok := cmd.Context().Value(constant.ClientContextKey).(*clickup.Client)

	if !ok {
		return fmt.Errorf("failed to get client")
	}

	spaces, _, err := client.Spaces.GetSpaces(cmd.Context(), teamID, archived)

	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(spaces, "", "  ")

	if err != nil {
		return err
	}

	cmd.Printf("%s\n", data)

	return nil
}

func init() {
	Command.AddCommand(listCommand)

	listCommand.PersistentFlags().BoolP("archived", "a", false, "archived")
}
