package component

import (
	"PROJECT_NAME/config"
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/johnpoint/go-bootstrap/core"
	ginBoot "github.com/johnpoint/go-bootstrap/gin"
	ginBootMiddlware "github.com/johnpoint/go-bootstrap/gin/middleware"
	"time"
)

type Api struct{}

var _ core.Component = (*Api)(nil)

func (r *Api) Init(c context.Context) error {
	return ginBoot.NewApiServer(config.Config.ApiServerListen,
		gin.Recovery(),
		cors.New(cors.Config{
			AllowOrigins:     config.Config.CORS,
			AllowMethods:     []string{"PUT", "GET", "POST", "DELETE"},
			AllowHeaders:     []string{"Origin", "content-type", "Cookie", "x-token"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}),
		ginBootMiddlware.LogPlusMiddleware(),
	).Init(c)
}
