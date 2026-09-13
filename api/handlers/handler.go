package handlers

import (
	"net/http"
	"yul-tracker/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *service.Service
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{Service: svc}
}

func (h *Handler) Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
