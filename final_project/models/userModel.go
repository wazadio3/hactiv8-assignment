package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `json:"username" gorm:"unique" validate:"required"`
	Email    string `json:"email" gorm:"unique" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Age      int    `json:"age" validate:"required,gt=8"`
}
