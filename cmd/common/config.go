package common

import (
	"context"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type (
	CTX_CONFIG_FILE struct{}
)

func ConfigViper(cfgFile string) string {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		fullpath := filepath.Join(home, ".config", "gzk9000")
		if _, err := os.Stat(fullpath); os.IsNotExist(err) {
			os.MkdirAll(fullpath, 0755)
		}

		viper.AddConfigPath(fullpath)
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")

		cfgFile = filepath.Join(fullpath, "config.yaml")
	}

	viper.AutomaticEnv()

	return cfgFile
}

func LoadConfigByContext(ctx context.Context) {
	cfgFile := ctx.Value(CTX_CONFIG_FILE{}).(string)
	cfgFile = ConfigViper(cfgFile)
	if err := viper.ReadInConfig(); err != nil {
		slog.Error("failed to read config", "error", err, "config", cfgFile)
		panic(err)
	}
}
