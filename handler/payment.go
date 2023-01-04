package handler

import (
	"io"
	"net/http"
	"project-go/models/payment"
	broadcast "project-go/utils/broadcaster"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type paymentHandler struct {
	paymentService payment.Service
}

func MigrateAndGetPayment(db *gorm.DB) (paymentHandler *paymentHandler) {
	db.AutoMigrate(&payment.Payment{})

	paymentRepository := payment.NewRepository(db)
	paymentService := payment.NewService(paymentRepository)
	paymentHandler = NewPaymentHandler(paymentService)
	return
}

func NewPaymentHandler(paymentService payment.Service) *paymentHandler {
	return &paymentHandler{paymentService}
}

// Create godoc
//
//	@Summary	Create a payment
//	@Tags		payment
//	@Accept		json
//	@Produce	json
//	@Param		paymentInput	body		payment.InputPayment	true	"Message body"
//	@Success	201				{object}	payment.Payment
//	@Failure	400				{object}	Response
//	@Failure	401				{object}	Response
//	@Router		/protected/payment [post]
//
//	@Security	BearerAuth
func (th *paymentHandler) Create(c *gin.Context) {
	// Get json body
	var input payment.InputPayment
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

	newPayment, err := th.paymentService.Create(input)
	if err != nil {
		response := &Response{
			Success: false,
			Message: "Error: Something went wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	broadcaster := broadcast.GetBroadcaster()
	broadcaster.Submit(newPayment)

	response := &Response{
		Success: true,
		Message: "New payment created",
		Data:    newPayment,
	}
	c.JSON(http.StatusCreated, response)
}

// GetAll godoc
//
//	@Summary	Get all payment
//	@Tags		payment
//	@Accept		json
//	@Produce	json
//	@Success	200	{array}		payment.Payment
//	@Failure	400	{object}	Response
//	@Failure	401	{object}	Response
//	@Router		/protected/payment [get]
//
//	@Security	BearerAuth
func (th *paymentHandler) GetAll(c *gin.Context) {
	payments, err := th.paymentService.GetAll()
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
		Data:    payments,
	})
}

func (th *paymentHandler) GetAllByStream(c *gin.Context) {
	var listener = make(chan interface{})
	broadcaster := broadcast.GetBroadcaster()
	broadcaster.Register(listener)

	defer broadcaster.Unregister(listener)

	clientGone := c.Request.Context().Done()
	c.Stream(func(w io.Writer) bool {
		select {
		case <-clientGone:
			return false
		case message := <-listener:
			c.SSEvent("message", message)
			return true
		}
	})
}

// GetById godoc
//
//	@Summary	Get payment by ID
//	@Tags		payment
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"Payment ID"
//	@Success	200	{object}	Response
//	@Failure	400	{object}	Response
//	@Failure	401	{object}	Response
//	@Router		/protected/payment/{id} [get]
//
//	@Security	BearerAuth
func (th *paymentHandler) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Message: "Wrong id parameter",
			Data:    err.Error(),
		})
		return
	}

	payment, err := th.paymentService.GetById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Message: "Payment not found",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		Success: true,
		Data:    payment,
	})
}

// Update by id godoc
//
//	@Summary	Update payment by ID
//	@Tags		payment
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"Payment ID"
//	@Success	201	{object}	Response
//	@Failure	400	{object}	Response
//	@Failure	401	{object}	Response
//	@Router		/protected/payment/{id} [put]
//
//	@Security	BearerAuth
func (th *paymentHandler) Update(c *gin.Context) {
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
	var input payment.InputPayment
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

	uPayment, err := th.paymentService.Update(id, input)
	if err != nil {
		response := &Response{
			Success: false,
			Message: "Error: Something went wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	broadcaster := broadcast.GetBroadcaster()
	broadcaster.Submit(uPayment)

	response := &Response{
		Success: true,
		Message: "payment updated",
		Data:    uPayment,
	}
	c.JSON(http.StatusCreated, response)
}

// Delete by id godoc
//
//	@Summary	Delete payment by ID
//	@Tags		payment
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"Payment ID"
//	@Success	200	{object}	Response
//	@Failure	400	{object}	Response
//	@Failure	401	{object}	Response
//	@Router		/protected/payment/{id} [delete]
//
//	@Security	BearerAuth
func (th *paymentHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Message: "Wrong id parameter",
			Data:    err.Error(),
		})
		return
	}

	err = th.paymentService.Delete(id)
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
		Message: "Payment successfully deleted",
	})
}
