package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ItemCode    string `json:"itemCode" gorm:"type:varchar(100)"`
	Description string `json:"description" gorm:"type:varchar(100)"`
	Quantity    int    `json:"quantity" gorm:"type:integer"`
	OrderID     uint   `json:"orderId" gorm:"type:integer"`
}
