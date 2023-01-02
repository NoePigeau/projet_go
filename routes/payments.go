package routes

import (
	"project-go/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initPaymentRoutes(public *gin.RouterGroup, protected *gin.RouterGroup, db *gorm.DB) {

	paymentHandler := handler.MigrateAndGetPayment(db)

	api := protected.Group("/payment")
	stream := public.Group("/stream")
	stream.GET("/", paymentHandler.GetAllByStream)

	api.POST("/", paymentHandler.Create)
	api.GET("/", paymentHandler.GetAll)
	api.GET("/:id", paymentHandler.GetById)
	api.PUT("/:id", paymentHandler.Update)
	api.DELETE("/:id", paymentHandler.Delete)
}
