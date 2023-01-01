package user

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
