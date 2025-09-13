package router

import (
	"github.com/facedamon/golang-homework/blog/handler"
	"github.com/facedamon/golang-homework/blog/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	g := gin.Default()
	g.Use(middleware.ErrorHandler())
	auth := g.Group("/api/auth")
	{
		auth.POST("/login", handler.Login)
		auth.POST("/register", handler.Register)
	}
	//需要验证的接口
	api := g.Group("/api/bus")
	api.Use(middleware.AuthMiddleWare())
	{
		api.POST("/createPost", handler.CreatePost)
		api.GET("/getAllPostList", handler.GetAllPostList)
		api.GET("/getPostById/:id", handler.GetPostById)
		api.POST("/updatePost", handler.UpdatePost)
		api.GET("/deletePostById/:id", handler.DeletePostById)

		api.POST("/createComment", handler.CreateComment)
		api.GET("/getCommentsByPostId/:id", handler.GetCommentsByPostId)
	}
	return g
}
