package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderedAt    time.Time `json:"orderedAt"`
	CustomerName string    `json:"customerName" gorm:"type:varchar(100)"`
}
