package handlers

import (
	"github.com/gin-gonic/gin"
)

type PingHandler struct {
}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

func (h *PingHandler) GetPing(c *gin.Context) {
	c.String(200, "pong")
}
