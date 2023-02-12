package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null"`
	Email    string `gorm:"type:varchar(20);not null"`
	Password string `gorm:"type:varchar(20);not null"`
}
