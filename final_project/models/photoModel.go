package models

import "gorm.io/gorm"

type Photo struct {
	gorm.Model
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" validate:"required"`
	UserID   uint   `json:"user_id"`
}
