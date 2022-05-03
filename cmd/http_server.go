package cmd

import (
	"PROJECT_NAME/depend"
	"PROJECT_NAME/pkg/bootstrap"
	"PROJECT_NAME/pkg/log"
	"context"
	"github.com/spf13/cobra"
)

var httpServerCommand = &cobra.Command{
	Use:   "api",
	Short: "Start http server",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		i := bootstrap.Helper{}
		i.AddComponent(
			&depend.Logger{},
			&depend.Api{},
		).SetLogger(log.GetLogger())
		err := i.Init(ctx)
		if err != nil {
			panic(err)
			return
		}

		forever := make(chan struct{})
		<-forever
	},
}
