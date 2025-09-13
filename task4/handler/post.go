package handler

import (
	"net/http"
	"strconv"

	"github.com/facedamon/golang-homework/blog/global"
	"github.com/facedamon/golang-homework/blog/model"
	"github.com/gin-gonic/gin"
)

func CreatePost(ctx *gin.Context) {
	global.Logger.Debug("开始创建文章...")
	value, _ := ctx.Get("user")
	user, _ := value.(*model.User)

	var post *model.Post
	if err := ctx.ShouldBindBodyWithJSON(&post); err != nil {
		global.Logger.Errorln("文章参数绑定失败", err)
		ctx.JSON(http.StatusOK, global.R{
			Code: http.StatusBadRequest,
			Msg:  "文章参数绑定失败",
			Data: err.Error(),
		})
		return
	}
	post.UserID = user.ID
	if err := global.Db.Create(post).Error; err != nil {
		global.Logger.Errorf("文章:%s,创建失败,%v\n", post.Title, err)
		ctx.JSON(http.StatusOK, global.R{
			Code: http.StatusInternalServerError,
			Data: err.Error(),
		})
		return
	}
	global.Logger.Debugf("文章:%s,创建成功\n", post.Title)

	ctx.JSON(http.StatusOK, global.R{
		Code: http.StatusOK,
	})
}

func GetAllPostList(ctx *gin.Context) {
	var posts *[]model.Post
	if err := global.Db.Debug().Find(&posts).Error; err != nil {
		global.Logger.Debug("GetAllPostList失败", err)
		ctx.JSON(http.StatusOK, global.R{
			Code: http.StatusInternalServerError,
			Msg:  "GetAllPostList失败",
			Data: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, global.R{
		Code: http.StatusOK,
		Data: posts,
	})
}

func GetPostById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusOK, global.R{
			Code: http.StatusOK,
		})
		return
	}
	post := &model.Post{}
	u, err := strconv.ParseUint(id, 10, 10)
	if err != nil {
		global.Logger.Debug("GetPostById中的id应为数字类型", err)
		ctx.JSON(http.StatusOK, global.R{
			Code: http.StatusInternalServerError,
			Msg:  "GetPostById中的id应为数字类型",
			Data: err.Error(),
		})
		return
	}
	post.ID = uint(u)
	if err := global.Db.Debug().Where("id=?", post.ID).Preload("User").First(&post).Error; err != nil {
		global.Logger.Debug("查询GetPostById失败", err)
		ctx.JSON(http.StatusOK, global.R{
			Code: http.StatusInternalServerError,
			Msg:  "查询GetPostById失败",
			Data: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, global.R{
		Code: http.StatusOK,
		Data: post,
	})

}

func UpdatePost(ctx *gin.Context) {
	value, _ := ctx.Get("user")
	user, _ := value.(*model.User)
	var input *model.Post
	if err := ctx.ShouldBindBodyWithJSON(&input); err != nil {
		global.Logger.Errorln("文章参数绑定失败", err)
		ctx.JSON(http.StatusOK, global.R{
			Code: http.StatusBadRequest,
			Msg:  "文章参数绑定失败",
			Data: err.Error(),
		})
		return
	}
	//更新时id不能为空
	if input.ID == 0 {
		global.Logger.Errorln("文章id不能为空")
		ctx.JSON(http.StatusOK, global.R{
			Code: http.StatusBadRequest,
			Msg:  "文章id不能为空",
		})
		return
	}
	//查询文章内容及所属人
	var post *model.Post
	if err := global.Db.Debug().Where("id=?", input.ID).Preload("User").Find(&post).Error; err != nil {
		global.Logger.Errorln("文章不存在")
		ctx.JSON(http.StatusOK, global.R{
			Code: http.StatusBadRequest,
			Msg:  "文章不存在",
		})
		return
	}
	if user.ID != post.UserID {
		global.Logger.Errorln("文章所属人与登陆人不一致,不允许修改")
		ctx.JSON(http.StatusOK, global.R{
			Code: http.StatusBadRequest,
			Msg:  "文章所属人与登陆人不一致",
		})
		return
	}
	//更新文章
	post.Title = input.Title
	post.Content = input.Content
	global.Db.Debug().Save(&post)
	ctx.JSON(http.StatusOK, global.R{
		Code: http.StatusOK,
	})
}

func DeletePostById(ctx *gin.Context) {
	value, _ := ctx.Get("user")
	user, _ := value.(*model.User)

	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusOK, global.R{
			Code: http.StatusOK,
		})
		return
	}
	u, err := strconv.ParseUint(id, 10, 10)
	if err != nil {
		global.Logger.Debug("DeletePostById中的id应为数字类型", err)
		ctx.JSON(http.StatusOK, global.R{
			Code: http.StatusInternalServerError,
			Msg:  "DeletePostById中的id应为数字类型",
			Data: err.Error(),
		})
		return
	}
	post := new(model.Post)
	post.ID = uint(u)
	//查询文章内容及所属人
	if err := global.Db.Debug().Where("id=?", post.ID).Preload("User").Find(&post).Error; err != nil {
		global.Logger.Errorln("文章不存在")
		ctx.JSON(http.StatusOK, global.R{
			Code: http.StatusBadRequest,
			Msg:  "文章不存在",
		})
		return
	}
	if user.ID != post.UserID {
		global.Logger.Errorln("文章所属人与登陆人不一致,不允许删除")
		ctx.JSON(http.StatusOK, global.R{
			Code: http.StatusBadRequest,
			Msg:  "文章所属人与登陆人不一致,不允许删除",
		})
		return
	}
	//删除文章
	if err := global.Db.Debug().Delete(&post).Error; err != nil {
		global.Logger.Errorln("删除文章失败", err)
		ctx.JSON(http.StatusOK, global.R{
			Code: http.StatusInternalServerError,
			Msg:  "删除文章失败",
			Data: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, global.R{
		Code: http.StatusOK,
	})
}
