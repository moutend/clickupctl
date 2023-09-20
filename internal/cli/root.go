package cli

import (
	"context"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/moutend/clickupctl/internal/cache"
	"github.com/moutend/clickupctl/internal/cli/checklist"
	"github.com/moutend/clickupctl/internal/cli/folder"
	"github.com/moutend/clickupctl/internal/cli/list"
	"github.com/moutend/clickupctl/internal/cli/space"
	"github.com/moutend/clickupctl/internal/cli/task"
	"github.com/moutend/clickupctl/internal/cli/team"
	"github.com/moutend/clickupctl/internal/cli/whoami"
	"github.com/moutend/clickupctl/internal/constant"
	"github.com/raksul/go-clickup/clickup"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCommand = &cobra.Command{
	Use:               "clickupctl",
	Short:             "command line client for clickup",
	SilenceUsage:      true,
	PersistentPreRunE: rootCommandPersistentPreRunE,
}

func rootCommandPersistentPreRunE(cmd *cobra.Command, args []string) error {
	viper.SetEnvPrefix("clickup")
	viper.BindEnv("api_token")

	transport, err := cache.NewTransport(cmd.Context())

	if err != nil {
		return err
	}
	if debug, _ := cmd.Flags().GetBool("debug"); debug {
		transport.SetLogger(log.New(cmd.OutOrStdout(), "Debug: ", 0))
	}

	httpClient := &http.Client{
		Transport: transport,
	}

	client := clickup.NewClient(httpClient, viper.GetString("api_token"))

	ctx := cmd.Context()
	ctx = context.WithValue(ctx, constant.TransportContextKey, transport)
	ctx = context.WithValue(ctx, constant.ClientContextKey, client)

	cmd.SetContext(ctx)

	return nil
}

func init() {
	RootCommand.PersistentFlags().BoolP("debug", "d", false, "enable debug output")

	if info, ok := debug.ReadBuildInfo(); ok {
		RootCommand.Version = info.Main.Version
	} else {
		RootCommand.Version = "undefined"
	}

	RootCommand.AddCommand(checklist.Command)
	RootCommand.AddCommand(folder.Command)
	RootCommand.AddCommand(list.Command)
	RootCommand.AddCommand(space.Command)
	RootCommand.AddCommand(task.Command)
	RootCommand.AddCommand(team.Command)
	RootCommand.AddCommand(whoami.Command)
}
