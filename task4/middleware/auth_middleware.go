package middleware

import (
	"net/http"

	"github.com/facedamon/golang-homework/blog/global"
	"github.com/facedamon/golang-homework/blog/model"
	"github.com/facedamon/golang-homework/blog/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusOK, global.R{
				Code: http.StatusUnauthorized,
				Msg:  "未授权",
			})
			ctx.Abort()
			return
		}
		username, err := utils.ParseJWT(token)
		if err != nil {
			ctx.JSON(http.StatusOK, global.R{
				Code: http.StatusUnauthorized,
				Msg:  "未授权",
				Data: err.Error(),
			})
			ctx.Abort()
			return
		}
		global.Logger.Debug("username=", username)
		var user *model.User
		if err := global.Db.Debug().Where("username=?", username).Find(&user).Error; err != nil {
			global.Logger.Debugf("用户%s不存在", username)
			ctx.JSON(http.StatusOK, global.R{
				Code: http.StatusUnauthorized,
				Msg:  "未授权",
			})
			ctx.Abort()
			return
		}
		ctx.Set("user", user)
		ctx.Next()
	}
}
