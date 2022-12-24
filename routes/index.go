package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoutes(r *gin.Engine, db *gorm.DB) {
	GetProductRoutes(r, db)
}
