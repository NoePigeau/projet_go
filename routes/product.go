package routes

import (
	"project-go/handler"
	"project-go/product"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initProductRoutes(r *gin.Engine, db *gorm.DB) {

	db.AutoMigrate(&product.Product{})

	api := r.Group("/product")
	productRepository := product.NewRepository(db)
	productService := product.NewService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	api.POST("/", productHandler.Create)
	api.GET("/", productHandler.GetAll)
	api.GET("/:id", productHandler.GetById)
	api.PUT("/:id", productHandler.Update)
	api.DELETE("/:id", productHandler.Delete)
}
