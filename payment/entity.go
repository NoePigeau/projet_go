package payment

import (
	"project-go/product"
	"time"
)

type Payment struct {
	ID        int              `json:"id" gorm:"primaryKey;autoIncrement:true"`
	PricePaid float32          `json:"price_paid"`
	ProductID int              `json:"payment_id"`
	Product   *product.Product `json:"product" gorm:"foreignKey:ProductID"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}
