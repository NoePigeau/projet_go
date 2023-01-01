package product

type InputProduct struct {
	Name  string  `json:"name" binding:"required"`
	Price float32 `json:"price" binding:"required"`
}
