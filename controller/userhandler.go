package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	// 获取参数
	// name := ctx.PostForm("name")
	phone := ctx.PostForm("phone")
	password := ctx.PostForm("password")
	if len(phone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code":422, "msg":"The phone num must be 11 digits!"})
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code":422, "msg":"Password can not be less than 6 digits!"})
	}
 	ctx.JSON(200, gin.H{
		"msg":"Register success",
	})
}