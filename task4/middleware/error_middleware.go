package middleware

import (
	"errors"
	"net/http"

	"github.com/facedamon/golang-homework/blog/global"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		for _, e := range ctx.Errors {
			err := e.Err
			var myErr *global.R
			if errors.As(err, &myErr) {
				ctx.JSON(http.StatusOK, myErr)
			} else {
				ctx.JSON(http.StatusOK, global.R{
					Code: http.StatusInternalServerError,
					Msg:  "服务器异常",
					Data: err.Error(),
				})
			}
			return
		}
	}
}
