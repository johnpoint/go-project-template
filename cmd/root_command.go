package cmd

import (
	"PROJECT_NAME/depend"
	"PROJECT_NAME/pkg/bootstrap"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

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
	cobra.OnInitialize(func() {
		if configPath == "" {
			configPath = "config_local.yaml"
		}
		bootstrap.AddGlobalComponent(
			&depend.Config{
				Path: configPath,
			},
			&depend.Logger{},
		)
	})
	rootCmd.PersistentFlags().StringVar(&configPath, "config", "config_local.json", "config file (default is ./config_local.yaml)")

	rootCmd.AddCommand(httpServerCommand)
}
