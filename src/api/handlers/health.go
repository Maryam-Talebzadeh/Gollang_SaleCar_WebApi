package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Workingggggggggggg")
	return
}

func (h *HealthHandler) HealthPost(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Working post")
	return
}

func (h *HealthHandler) HealthGetById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Working get by Id")
	return
}
