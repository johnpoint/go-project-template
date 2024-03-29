// Code generated by newapi. DO NOT EDIT.
//
// command: newapi gen --config api.yaml
//
// StateEndpoint 获取API服务状态 GET /state 系统API State
package endpoints

import (
	"PROJECT_NAME/services"
	"github.com/gin-gonic/gin"
	ginBoot "github.com/johnpoint/go-bootstrap/gin"
)

type StateEndpoint struct {
	ginBoot.BaseEndpoint
	ginBoot.JsonCodecEp
	ginBoot.GETEndpoint
}

func (e *StateEndpoint) Path() string {
	return "/state"
}

func (e *StateEndpoint) HandlerFunc() gin.HandlerFunc {
	return ginBoot.Endpoint(e, services.StateService)
}

func init() {
	var endpoint StateEndpoint
	endpoint.SetMiddleware(services.StateSetMiddleware())
	if err := ginBoot.Server().AddEndpoint(&endpoint); err != nil {
		panic(err)
	}
}
