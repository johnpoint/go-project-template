package cmd

import (
	"PROJECT_NAME/depend"
	"PROJECT_NAME/pkg/log"
	"context"

	"github.com/johnpoint/go-bootstrap"
	"github.com/spf13/cobra"
)

var httpServerCommand = &cobra.Command{
	Use:   "api",
	Short: "Start http server",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		err := bootstrap.NewBoot(
			ctx,
			&depend.Logger{},
			&depend.Api{},
		).WithLogger(log.GetLogger()).Init()
		if err != nil {
			panic(err)
			return
		}

		forever := make(chan struct{})
		<-forever
	},
}
