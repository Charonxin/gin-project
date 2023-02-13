package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"main.go/common"
	"main.go/model"
	"main.go/response"
	"net/http"
)

func Register(ctx *gin.Context) {
	DB := common.GetDB()
	var requestUser = model.User{}
	ctx.BindJSON(&requestUser)
	// 获取参数
	username := requestUser.Username
	email := requestUser.Email
	password := requestUser.Password

	//数据验证
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	log.Println(username, email, password)

	if isUsernameExist(DB, username) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户名已存在")
		return
	}
	if isEmailExist(DB, email) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "邮箱已存在")
		return
	}

	// 创建用户
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}
	newUser := model.User{
		Username: username,
		Email:    email,
		Password: string(hasedPassword),
	}
	DB.Create(&newUser)

	// 发放token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		//response.Fail(ctx,"系统异常",gin.H{"code":500})
		log.Println("error" + err.Error())
		return
	}
	response.Success(ctx, gin.H{"token": token}, "注册成功")

}

func isUsernameExist(db *gorm.DB, username string) bool {
	var user model.User
	db.Where("username = ?", username).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func isEmailExist(db *gorm.DB, email string) bool {
	var user model.User
	db.Where("username = ?", email).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
