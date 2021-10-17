package cmd

import (
	"{{project_name}}/config"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{}
	cfgFile string
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
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config_local.json", "config file (default is ./config_local.json)")

	rootCmd.AddCommand(genConfigCommand)
	rootCmd.AddCommand(httpServerCommand)
}

func initConfig() {
	config.Config = &config.ServiceConfig{ConfigFile: cfgFile}
}
