package list

import (
	"encoding/json"
	"fmt"
	"strconv"

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

	id, err := strconv.ParseInt(args[0], 10, 64)

	if err != nil {
		return err
	}

	var list clickup.List

	if folderless, _ := cmd.Flags().GetBool("folderless"); folderless {
		spaceID := int(id)
		request := buildListRequest(cmd)

		list, _, err = client.Lists.CreateFolderlessList(cmd.Context(), spaceID, request)
	}
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
	Command.AddCommand(createCommand)

	createCommand.PersistentFlags().BoolP("folderless", "f", false, "folderless")

	setupListRequestFlags(createCommand)
}
