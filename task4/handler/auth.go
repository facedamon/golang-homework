package handler

import (
	"net/http"

	"github.com/facedamon/golang-homework/blog/global"
	"github.com/facedamon/golang-homework/blog/model"
	"github.com/facedamon/golang-homework/blog/utils"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var input struct {
		Username string `json:"Username"`
		Password string `json:"Password"`
	}

	if err := ctx.ShouldBindBodyWithJSON(&input); err != nil {
		global.Logger.Errorln("登录参数绑定失败", err)
		ctx.AbortWithStatusJSON(http.StatusOK, global.R{
			Code: http.StatusBadRequest,
			Msg:  "登录参数绑定失败",
			Data: err,
		})
		return
	}

	user := model.User{}
	//判断用户是否存在
	if err := global.Db.Debug().Where("Username=?", input.Username).Find(&user).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, global.R{
			Code: http.StatusUnauthorized,
			Msg:  "未授权",
			Data: err,
		})
		return
	}
	global.Logger.Debug(user)
	if !utils.CheckPasswd(input.Password, user.Password) {
		p, _ := utils.HashPasswd(input.Password)
		global.Logger.Errorf("密码错误。明文:%s, 密文:%s, 真实:%s", input.Password, p, user.Password)
		ctx.AbortWithStatusJSON(http.StatusOK, global.R{
			Code: http.StatusUnauthorized,
			Msg:  "密码错误",
		})
		return
	}

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		global.Logger.Errorln("生成token失败", err)
		ctx.JSON(http.StatusOK, global.R{
			Code: http.StatusInternalServerError,
			Msg:  "生成token失败",
			Data: err,
		})
		return
	}

	//user1: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTc3NDE2MDIsInVzZXJuYW1lIjoidG9tIn0.G0YGmrpFN-TudBf6ueA0uTKpncYVVdsr3wyoO0B5wFo
	//user2: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTc3NTY5OTksInVzZXJuYW1lIjoid3N4In0.sIwEUy0Fsbx8SkmtI6OE_Y8oKWYe16x3Y6BO9RkxbQk
	ctx.AbortWithStatusJSON(http.StatusOK, global.R{
		Code: http.StatusOK,
		Data: token,
	})
}

func Register(ctx *gin.Context) {
	var input struct {
		Username string
		Password string
		Email    string
	}
	if err := ctx.ShouldBindBodyWithJSON(&input); err != nil {
		global.Logger.Errorln("注册参数绑定失败", err)
		ctx.JSON(http.StatusOK, global.R{
			Code: http.StatusBadRequest,
			Msg:  "注册参数绑定失败",
			Data: err,
		})
		return
	}
	hashPwd, err := utils.HashPasswd(input.Password)
	global.Logger.Debugf("注册时明文:%s, 加密后:%s", input.Password, hashPwd)
	if err != nil {
		global.Logger.Errorln("加密异常", err)
		ctx.JSON(http.StatusOK, global.R{
			Code: http.StatusInternalServerError,
			Msg:  "加密异常",
			Data: err,
		})
		return
	}
	user := model.User{}
	user.Password = hashPwd
	user.Username = input.Username
	user.Email = input.Email

	if err = global.Db.Create(&user).Error; err != nil {
		global.Logger.Errorln("无法创建用户", err)
		ctx.JSON(http.StatusOK, global.R{
			Code: http.StatusInternalServerError,
			Msg:  "无法创建记录",
			Data: err,
		})
		return
	}

	ctx.JSON(http.StatusOK, global.R{
		Code: http.StatusOK,
	})
	return
}
