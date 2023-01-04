package handler

import (
	"net/http"
	"project-go/models/product"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type productHandler struct {
	productService product.Service
}

func MigrateAndGetProduct(db *gorm.DB) (productHandler *productHandler) {
	db.AutoMigrate(&product.Product{})

	productRepository := product.NewRepository(db)
	productService := product.NewService(productRepository)
	productHandler = NewProductHandler(productService)
	return
}

func NewProductHandler(productService product.Service) *productHandler {
	return &productHandler{productService}
}

// Create godoc
//
//	@Summary	Create a product
//	@Tags		product
//	@Accept		json
//	@Produce	json
//	@Param		productInput	body		product.InputProduct	true	"Product body"
//	@Success	201				{object}	product.Product
//	@Failure	400				{object}	Response
//	@Failure	401				{object}	Response
//	@Router		/protected/product [post]
//
//	@Security	BearerAuth
func (th *productHandler) Create(c *gin.Context) {
	// Get json body
	var input product.InputProduct
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := &Response{
			Success: false,
			Message: "Cannot extract JSON body",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newProduct, err := th.productService.Create(input)
	if err != nil {
		response := &Response{
			Success: false,
			Message: "Error: Something went wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := &Response{
		Success: true,
		Message: "New product created",
		Data:    newProduct,
	}
	c.JSON(http.StatusCreated, response)
}

// GetAll godoc
//
//	@Summary	Get all product
//	@Tags		product
//	@Accept		json
//	@Produce	json
//	@Success	200	{array}		Response
//	@Failure	400	{object}	Response
//	@Failure	401	{object}	Response
//	@Router		/protected/product [get]
//
//	@Security	BearerAuth
func (th *productHandler) GetAll(c *gin.Context) {
	products, err := th.productService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Message: "Error: Something went wrong",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		Success: true,
		Data:    products,
	})
}

// GetById godoc
//
//	@Summary	Get product by ID
//	@Tags		product
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"Product ID"
//	@Success	200	{object}	Response
//	@Failure	400	{object}	Response
//	@Failure	401	{object}	Response
//	@Router		/protected/product/{id} [get]
//
//	@Security	BearerAuth
func (th *productHandler) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Message: "Wrong id parameter",
			Data:    err.Error(),
		})
		return
	}

	product, err := th.productService.GetById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Message: "Product not found",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		Success: true,
		Data:    product,
	})
}

// Update godoc
//
//	@Summary	Update product by ID
//	@Tags		product
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"Product ID"
//	@Success	200	{object}	Response
//	@Failure	400	{object}	Response
//	@Failure	401	{object}	Response
//	@Router		/protected/product/{id} [put]
//
//	@Security	BearerAuth
func (th *productHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Message: "Wrong id parameter",
			Data:    err.Error(),
		})
		return
	}

	// Get json body
	var input product.InputProduct
	err = c.ShouldBindJSON(&input)
	if err != nil {
		response := &Response{
			Success: false,
			Message: "Cannot extract JSON body",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	uProduct, err := th.productService.Update(id, input)
	if err != nil {
		response := &Response{
			Success: false,
			Message: "Product not found",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := &Response{
		Success: true,
		Message: "Product updated",
		Data:    uProduct,
	}
	c.JSON(http.StatusCreated, response)
}

// Delete godoc
//
//	@Summary	Delete product by ID
//	@Tags		product
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"Product ID"
//	@Success	200	{object}	Response
//	@Failure	400	{object}	Response
//	@Failure	401	{object}	Response
//	@Router		/protected/product/{id} [delete]
//
//	@Security	BearerAuth
func (th *productHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Message: "Wrong id parameter",
			Data:    err.Error(),
		})
		return
	}

	err = th.productService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Message: "Error: Something went wrong",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		Success: true,
		Message: "Product successfully deleted",
	})
}
