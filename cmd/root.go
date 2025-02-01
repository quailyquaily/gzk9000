package cmd

import (
	"context"
	"log/slog"
	"os"

	"github.com/quailyquaily/gzk9000/cmd/common"
	"github.com/quailyquaily/gzk9000/cmd/gen"
	"github.com/quailyquaily/gzk9000/cmd/migrate"
	"github.com/quailyquaily/gzk9000/cmd/server"
	"github.com/quailyquaily/gzk9000/store"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
)

// rootCmd is the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gzk9000",
	Short: "gzk9000 is an API proxy cluster tool.",
	Long:  `gzk9000 is an API proxy cluster tool that enables multiple proxies to bypass IP-based rate limits across different servers.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		ctx = context.WithValue(ctx, common.CTX_CONFIG_FILE{}, cfgFile)

		common.LoadConfigByContext(ctx)
		h := store.MustInit(store.Config{
			Driver: viper.GetString("db.driver"),
			DSN:    viper.GetString("db.dsn"),
		})

		ctx = store.NewContext(ctx, h)

		cmd.SetContext(ctx)
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func ExecuteContext(ctx context.Context) error {
	return rootCmd.ExecuteContext(ctx)
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.config/gzk9000/config.yaml)")

	// Add subcommands here
	rootCmd.AddCommand(server.NewCmd())
	rootCmd.AddCommand(gen.NewCmd())
	rootCmd.AddCommand(migrate.NewCmd())
}

func initConfig() {
	cfgFile = common.ConfigViper(cfgFile)
	if err := viper.ReadInConfig(); err != nil {
		// if the config file does not exist, ask the user to login
		slog.Info("Config file does not exist. Please create one\ndefault is $HOME/.config/gzk9000/config.yaml, or specify with --config")
		os.Exit(1)
	}
	slog.Info("Config file loaded", "config", cfgFile)
}
