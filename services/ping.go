package services

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

type PingRequest struct {
	Message string `json:"message" form:"message"`
}

type PingResponse struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

func PingSetMiddleware() []gin.HandlerFunc {
	return []gin.HandlerFunc{}
}

func PingService(ctx context.Context, req *PingRequest) (resp *PingResponse, err error) {
	return &PingResponse{
		Message: req.Message,
		Time:    time.Now(),
	}, nil
}
