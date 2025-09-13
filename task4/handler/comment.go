package handler

import (
	"net/http"
	"strconv"

	"github.com/facedamon/golang-homework/blog/global"
	"github.com/facedamon/golang-homework/blog/model"
	"github.com/gin-gonic/gin"
)

func CreateComment(ctx *gin.Context) {
	global.Logger.Debug("开始添加评论...")
	value, _ := ctx.Get("user")
	user, _ := value.(*model.User)

	var input struct {
		Content string `json:"content" binding:"required" msg:"评论内容不能为空"`
		UserID  uint
		PostID  uint `json:"postId" binding:"required" msg:"评论文章id不能为空"`
	}
	if err := ctx.ShouldBindBodyWithJSON(&input); err != nil {
		global.Logger.Errorln("评论参数绑定失败", err)
		ctx.JSON(http.StatusOK, global.R{
			Code: http.StatusBadRequest,
			Msg:  "评论参数绑定失败",
			Data: err.Error(),
		})
		return
	}
	comment := new(model.Comment)
	comment.PostID = input.PostID
	comment.UserID = user.ID
	comment.Content = input.Content
	if err := global.Db.Create(comment).Error; err != nil {
		global.Logger.Errorln("评论创建失败", err)
		ctx.JSON(http.StatusOK, global.R{
			Code: http.StatusInternalServerError,
			Data: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, global.R{
		Code: http.StatusOK,
	})
}

func GetCommentsByPostId(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusOK, global.R{
			Code: http.StatusOK,
		})
		return
	}
	u, err := strconv.ParseUint(id, 10, 10)
	if err != nil {
		global.Logger.Errorln("GetCommentsByPostId中的id应为数字类型", err)
		ctx.JSON(http.StatusOK, global.R{
			Code: http.StatusInternalServerError,
			Msg:  "GetCommentsByPostId中的id应为数字类型",
			Data: err.Error(),
		})
		return
	}
	postId := uint(u)
	comment := &[]model.Comment{}
	if err := global.Db.Debug().Where("post_id=?", postId).Find(&comment).Error; err != nil {
		global.Logger.Errorf("查询文章id:%s的评论列表失败, err:=%v\n", postId, err)
		ctx.JSON(http.StatusOK, global.R{
			Code: http.StatusInternalServerError,
			Msg:  "查询文章评论列表失败",
			Data: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, global.R{
		Code: http.StatusOK,
		Data: comment,
	})
}
