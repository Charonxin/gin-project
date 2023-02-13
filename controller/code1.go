package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"main.go/common"
	"net/http"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null"`
	Email    string `gorm:"type:varchar(20);not null"`
	Password string `gorm:"type:varchar(20);not null"`
}

type register struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Confirm  string `json:"confirm"`
}

func IsExistUsername(db *gorm.DB, username string, email string) bool {
	var user User
	db.Where("username = ?", username).First(&user)
	if user.ID != 0 {
		return true
	}
	db.Where("email = ?", email).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func Regis(context *gin.Context) {
	db := common.GetDB()
	register := register{}
	err := context.BindJSON(&register)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(register)

	if IsExistUsername(db, register.Username, register.Email) {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已存在"})
		return
	}

	newUser := User{
		Username: register.Username,
		Email:    register.Email,
		Password: register.Password,
	}
	db.Create(&newUser)
	log.Println("success")
}
