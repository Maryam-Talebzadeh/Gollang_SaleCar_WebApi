package routers

import (
	"github.com/Maryam-Talebzadeh/Gollang_SaleCar_WebApi/src/api/handlers"
	"github.com/gin-gonic/gin"
)

func Health(r *gin.RouterGroup) {
	handler := handlers.NewHealthHandler()

	r.GET("/", handler.Health)
	r.POST("/", handler.HealthPost)
	r.GET("/:id", handler.HealthGetById)
}
