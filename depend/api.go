package depend

import (
	"PROJECT_NAME/app/controller"
	"PROJECT_NAME/config"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/johnpoint/go-bootstrap"
)

type Api struct{}

var _ bootstrap.Component = (*Api)(nil)

func (r *Api) Init(ctx context.Context) error {
	gin.SetMode(gin.ReleaseMode)
	routerGin := gin.New()
	routerGin.GET("/ping", controller.Pong)

	go func() {
		fmt.Println("[init] HTTP Listen at " + config.Config.HttpServerListen)
		err := routerGin.Run(config.Config.HttpServerListen)
		if err != nil {
			panic(err)
		}
	}()
	return nil
}
