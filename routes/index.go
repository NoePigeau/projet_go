package routes

import (
	"project-go/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(r *gin.Engine, db *gorm.DB) {

	public := r.Group("/api/public")
	initUserRoutes(public, db)

	protected := r.Group("/api/protected")
	protected.Use(middlewares.JwtAuthMiddleware())
	initProductRoutes(protected, db)
	initPaymentRoutes(protected, db)

}
