package payment

type InputPayment struct {
	ProductID int     `json:"product_id" binding:"required"`
	PricePaid float32 `json:"price_paid" binding:"required"`
}
