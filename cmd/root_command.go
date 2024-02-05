package cmd

import (
	"PROJECT_NAME/config"
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

func Level(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

var (
	rootCmd    = &cobra.Command{}
	configPath string
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&configPath, "config", "config_local.yaml", "config file (default is ./config_local.yaml)")

	rootCmd.AddCommand(httpServerCommand)
	config.InitConfig(configPath)
}
