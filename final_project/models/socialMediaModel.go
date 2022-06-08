package models

import "gorm.io/gorm"

type SosMed struct {
	gorm.Model
	Name           string `json:"name" validate:"required"`
	SocialMediaUrl string `json:"social_media_url" validate:"required"`
	UserID         uint   `json:"user_id"`
}
