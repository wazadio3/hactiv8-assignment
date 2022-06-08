package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserID  uint   `json:"user_id"`
	PhotoID uint   `json:"photo_id" validate:"required"`
	Message string `json:"message" validate:"required"`
}
