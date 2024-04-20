package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Username string `json:"username"`
	Role     string `json:"role"`
}
