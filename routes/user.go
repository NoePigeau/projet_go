package routes

import (
	"project-go/controllers/user"
	"project-go/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initUserRoutes(public *gin.RouterGroup, db *gorm.DB) {

	db.AutoMigrate(&user.User{})

	api := public.Group("/user")
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	api.POST("/register", userHandler.Register)
	api.POST("/login", userHandler.Login)
	api.GET("/current", userHandler.CurrentUser)
}
