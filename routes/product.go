package routes

import (
	"project-go/handler"
	"project-go/product"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetProductRoutes(r *gin.Engine, db *gorm.DB) {

	api := r.Group("/product")
	productRepository := product.NewRepository(db)
	productService := product.NewService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	api.POST("/", productHandler.Store)
	api.GET("/", productHandler.FetchAll)
	api.GET("/:id", productHandler.FetchById)
	api.PUT("/:id", productHandler.Update)
	api.DELETE("/:id", productHandler.Delete)
}
