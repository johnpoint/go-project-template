package cmd

import (
	"PROJECT_NAME/component"
	"context"

	"github.com/johnpoint/go-bootstrap/core"
	"github.com/johnpoint/go-bootstrap/log"
	"github.com/spf13/cobra"
)

var httpServerCommand = &cobra.Command{
	Use:   "api",
	Short: "Start http server",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		err := core.NewBoot(
			&component.Logger{},
			&component.Api{},
		).WithContext(ctx).WithLogger(log.GetLogger()).Init()
		if err != nil {
			panic(err)
		}

		forever := make(chan struct{})
		<-forever
	},
}
