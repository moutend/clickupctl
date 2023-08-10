package list

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

	client, ok := cmd.Context().Value(constant.ClientContextKey).(*clickup.Client)

	if !ok {
		return fmt.Errorf("failed to get client")
	}

	archived, _ := cmd.Flags().GetBool("archived")
	folderless, _ := cmd.Flags().GetBool("folderless")

	var lists []clickup.List
	var err error

	if folderless {
		spaceID := args[0]

		lists, _, err = client.Lists.GetFolderlessLists(cmd.Context(), spaceID, archived)
	} else {
		folderID := args[0]

		lists, _, err = client.Lists.GetLists(cmd.Context(), folderID, archived)
	}
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(lists, "", "  ")

	if err != nil {
		return err
	}

	cmd.Printf("%s\n", data)

	return nil
}

func init() {
	Command.AddCommand(listCommand)

	listCommand.PersistentFlags().BoolP("archived", "a", false, "archived")
	listCommand.PersistentFlags().BoolP("folderless", "f", false, "folderless")
}
