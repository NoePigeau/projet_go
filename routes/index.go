package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(r *gin.Engine, db *gorm.DB) {
	initProductRoutes(r, db)
	initPaymentRoutes(r, db)
}
