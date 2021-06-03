package models

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"gorm_model"`
	ID         uint   `gorm:"primary_key" form:"ID" json:"ID" binding:"required"`
	Login      string `form:"login" json:"login" binding:"required"`
	Name       string `form:"name" json:"name" binding:"required"`
	Phone      string `form:"phone" json:"phone" binding:"required"`
}

