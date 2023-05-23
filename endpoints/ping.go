package endpoints

import (
	"PROJECT_NAME/services"
	"github.com/gin-gonic/gin"
	ginBoot "github.com/johnpoint/go-bootstrap/gin"
)

type PingEndpoint struct {
	ginBoot.BaseEndpoint
	ginBoot.JsonCodecEp
	ginBoot.GETEndpoint
}

func (e *PingEndpoint) Path() string {
	return "/ping"
}

func (e *PingEndpoint) HandlerFunc() gin.HandlerFunc {
	return ginBoot.Endpoint(e, services.PingService)
}

func init() {
	var endpoint PingEndpoint
	endpoint.SetMiddleware(services.PingSetMiddleware())
	if err := ginBoot.Server().AddEndpoint(&endpoint); err != nil {
		panic(err)
	}
}
