package whoami

import (
	"encoding/json"
	"fmt"

	"github.com/moutend/clickupctl/internal/constant"
	"github.com/raksul/go-clickup/clickup"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:  "whoami",
	RunE: commandRunE,
}

func commandRunE(cmd *cobra.Command, args []string) error {
	client, ok := cmd.Context().Value(constant.ClientContextKey).(*clickup.Client)

	if !ok {
		return fmt.Errorf("failed to get client")
	}

	myself, _, err := client.Authorization.GetAuthorizedUser(cmd.Context())

	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(myself, "", "  ")

	if err != nil {
		return err
	}

	cmd.Printf("%s\n", data)

	return nil
}
