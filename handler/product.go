package handler

import (
	"net/http"
	"project-go/product"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type productHandler struct {
	productService product.Service
}

func NewProductHandler(productService product.Service) *productHandler {
	return &productHandler{productService}
}

func (ts *productHandler) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}

func (th *productHandler) Store(c *gin.Context) {
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

	newProduct, err := th.productService.Store(input)
	if err != nil {
		response := &Response{
			Success: false,
			Message: "Something went wrong",
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

func (th *productHandler) FetchAll(c *gin.Context) {
	products, err := th.productService.FetchAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		Success: true,
		Data:    products,
	})
}

func (th *productHandler) FetchById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Message: "Wrong id parameter",
			Data:    err.Error(),
		})
		return
	}

	product, err := th.productService.FetchById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		Success: true,
		Data:    product,
	})
}

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
			Message: "Something went wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := &Response{
		Success: true,
		Message: "New product created",
		Data:    uProduct,
	}
	c.JSON(http.StatusCreated, response)
}

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
			Message: "Something went wrong",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		Success: true,
		Message: "Product successfully deleted",
	})
}
