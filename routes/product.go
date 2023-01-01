package routes

import (
	"project-go/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initProductRoutes(protected *gin.RouterGroup, db *gorm.DB) {

	productHandler := handler.MigrateAndGetProduct(db)

	api := protected.Group("/product")
	api.POST("/", productHandler.Create)
	api.GET("/", productHandler.GetAll)
	api.GET("/:id", productHandler.GetById)
	api.PUT("/:id", productHandler.Update)
	api.DELETE("/:id", productHandler.Delete)
}
