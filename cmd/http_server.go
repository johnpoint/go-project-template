package cmd

import (
	"PROJECT_NAME/config"
	_ "PROJECT_NAME/endpoints"
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/johnpoint/go-bootstrap/v2/core"
	ginBoot "github.com/johnpoint/go-bootstrap/v2/gin"
	ginBootMiddlware "github.com/johnpoint/go-bootstrap/v2/gin/middleware"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var httpServerCommand = &cobra.Command{
	Use:   "api",
	Short: "Start http server",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		err := core.NewBoot(
			core.WithComponents(
				apiServer(config.Config),
			),
			core.Level(Level(config.Config.LogLevel)),
			core.LogOutput(os.Stderr),
			core.SetLoggerType(core.LoggerTypeJSON),
			core.WithContext(ctx),
		).Init()
		if err != nil {
			panic(err)
		}

		forever := make(chan struct{})
		<-forever
	},
}

var apiServer = func(config *config.ServiceConfig) *ginBoot.ApiServer {
	return ginBoot.NewApiServer(config.ApiServerListen,
		gin.Recovery(),
		cors.New(cors.Config{
			AllowOrigins:     config.CORS,
			AllowMethods:     []string{"PUT", "GET", "POST", "DELETE"},
			AllowHeaders:     []string{"Origin", "content-type", "Cookie", "x-token"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}),
		ginBootMiddlware.LogPlusMiddleware(),
	)
}
