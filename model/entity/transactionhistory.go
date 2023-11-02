package entity

import "time"

type TransactionHistory struct {
	ID         int `gorm:"primaryKey"`
	ProductID  int `gorm:"type:int"`
	UserID     int `gorm:"type:int"`
	Quantity   int `gorm:"type:int"`
	TotalPrice int `gorm:"type:int"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
