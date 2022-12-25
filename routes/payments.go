package routes

import (
	"project-go/handler"
	"project-go/payment"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initPaymentRoutes(r *gin.Engine, db *gorm.DB) {

	db.AutoMigrate(&payment.Payment{})

	api := r.Group("/payment")
	paymentRepository := payment.NewRepository(db)
	paymentService := payment.NewService(paymentRepository)
	paymentHandler := handler.NewPaymentHandler(paymentService)

	api.POST("/", paymentHandler.Create)
	api.GET("/", paymentHandler.GetAll)
	api.GET("/:id", paymentHandler.GetById)
	api.PUT("/:id", paymentHandler.Update)
	api.DELETE("/:id", paymentHandler.Delete)
}
