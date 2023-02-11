package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null"`
	Email    string `gorm:"type:varchar(20);not null"`
	Password string `gorm:"type:varchar(20);not null"`
}

func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "gindatabase1"
	username := "root"
	password := "qwer"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database, err " + err.Error())
	}
	db.AutoMigrate(&User{})
	return db
}

func GetDB() *gorm.DB {
	return db
}
