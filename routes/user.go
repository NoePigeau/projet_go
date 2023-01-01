package routes

import (
	"project-go/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initUserRoutes(public *gin.RouterGroup, protected *gin.RouterGroup, db *gorm.DB) {

	userHandler := handler.MigrateAndGetUser(db)

	api := public.Group("/user")
	api.POST("/register", userHandler.Register)
	api.POST("/login", userHandler.Login)

	api = protected.Group("/user")
	api.GET("/current", userHandler.CurrentUser)
}
