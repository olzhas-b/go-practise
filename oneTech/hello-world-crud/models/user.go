package models

import (
	"gorm.io/gorm"
)
type User struct {
	gorm.Model `json:"gorm_model"`
	ID	   	   uint   `gorm:"primaryKey"`
	Email      string
	Password   string
	Name       string
	Phone      string
}

//type Message struct {
//	gorm.Model
//	Title string
//	Time time.Time
//	UserID uint
//}

