package entity

import "time"

type Product struct {
	ID         int    `gorm:"primaryKey"`
	Title      string `gorm:"type:varchar(255)"`
	Price      int    `gorm:"type:int"`
	Stock      int    `gorm:"type:int"`
	CategoryID int    `gorm:"type:int"`
	CreateAt   time.Time
	UpdateAt   time.Time
}
