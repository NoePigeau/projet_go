package routes

import (
	"project-go/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(r *gin.Engine, db *gorm.DB) {

	public := r.Group("/api/public")
	protected := r.Group("/api/protected")
	protected.Use(middlewares.JwtAuthMiddleware())

	initUserRoutes(public, protected, db)
	initProductRoutes(protected, db)
	initPaymentRoutes(protected, db)
}
