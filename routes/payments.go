package routes

import (
	"project-go/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initPaymentRoutes(protected *gin.RouterGroup, db *gorm.DB) {

	paymentHandler := handler.MigrateAndGetPayment(db)

	api := protected.Group("/payment")
	api.POST("/", paymentHandler.Create)
	api.GET("/", paymentHandler.GetAll)
	api.GET("/sse", paymentHandler.GetAllByStream)
	api.GET("/:id", paymentHandler.GetById)
	api.PUT("/:id", paymentHandler.Update)
	api.DELETE("/:id", paymentHandler.Delete)
}
