package entity

import "time"

type Category struct {
	ID                int    `gorm:"primaryKey"`
	Type              string `gorm:"type:varchar(255)"`
	SoldProductAmount int    `gorm:"type:int"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
