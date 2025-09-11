package main

import (
	"log"

	"github.com/facedamon/golang-homework/gorm_test/config"
	"github.com/facedamon/golang-homework/gorm_test/global"
	"github.com/facedamon/golang-homework/gorm_test/model"
)

func main() {
	//gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/golang_test"), &gorm.Config{})
	config.InitConfig()
	log.Println(config.AppConfig.Database.Dsn)
	config.InitDB()

	//user := new(model.User)
	//post := new(model.Post)
	//comm := new(model.Comment)
	//
	//if err := global.Db.AutoMigrate(user, post, comm); err != nil {
	//	log.Fatalln("AutoMigrate failed.", err)
	//	return
	//}
	//log.Println("AutoMigrate successfully.")

	//加入模拟数据
	//user := &model.User{
	//	Name: "tom",
	//	Posts: []model.Post{
	//		{
	//			Title:   "我的世界",
	//			Content: "我的世界的内容",
	//		},
	//		{
	//			Title:   "新疆雪菊",
	//			Content: "新疆雪菊的功效",
	//		},
	//	},
	//}
	//global.Db.Save(user)
	//log.Println(user)

	//根据已有数据添加数据
	//user := &model.User{}
	//if err := global.Db.Debug().Preload("Posts").Where("name = ?", "tom").Find(user).Error; err != nil {
	//	log.Fatalln("查询就失败", err)
	//	return
	//}
	//bs, err := json.Marshal(user)
	//if err != nil {
	//	log.Fatalln("解析json出错", err)
	//	return
	//}
	//log.Println(string(bs))
	//var comms []model.Comment
	//for _, post := range user.Posts {
	//	comms = append(comms, model.Comment{
	//		PostID: post.ID,
	//		Comm:   "评论" + post.Title,
	//	})
	//}
	//if err = global.Db.Save(&comms).Error; err != nil {
	//	log.Fatalln("写入评论失败", err)
	//	return
	//}
	////读取某个用户发布的所有文章及其评论
	////多级预加载
	//if err = global.Db.Preload("Posts.Comments").Where("name=?", "tom").Find(&user).Error; err != nil {
	//	log.Fatalln("读取某个用户发布的所有文章及其评论失败", err)
	//	return
	//}
	//js, err := json.MarshalIndent(&user, "", " ")
	//if err != nil {
	//	log.Fatalln("json格式化错误", err)
	//	return
	//}
	//log.Println(string(js))

	//查询评论数量最多的文章
	//select post_id, comm, count(1) as cc  from comments c  group by post_id, comm order by cc desc limit 1
	//var comm model.Comment
	//if err := global.Db.Model(&model.Comment{}).Select("post_id, comm, count(1) as cc").Group("post_id").Group("comm").Order("cc desc").Limit(1).Scan(&comm).Error; err != nil {
	//	log.Fatalln("查询评论数量最多的文章id报错", err)
	//	return
	//}
	//log.Println(comm.PostID)
	//var posts []model.Post
	//if err := global.Db.Where("id=?", comm.PostID).Find(&posts).Error; err != nil {
	//	log.Fatalln("查询评论数量最多的文章信息报错", err)
	//	return
	//}
	//log.Println(posts)

	//在post创建时自动更新用户的文章数量统计字段
	//if err := global.Db.Debug().Create(&model.Post{UserID: 1, Title: "新的一天", Content: "新的一天的内容"}).Error; err != nil {
	//	log.Fatalln("在post创建时自动更新用户的文章数量统计字段失败", err)
	//	return
	//}
	//var user model.User
	//if err := global.Db.Debug().Model(&model.User{}).Where("id=?", 1).Find(&user).Error; err != nil {
	//	log.Fatalln("查询用户文章统计信息", err)
	//	return
	//}
	//log.Println(user.Name, *user.Num)

	//删除评论时，触发钩子函数,全部删除post_id=1的
	if err := global.Db.Debug().Where("post_id=?", 1).Delete(&model.Comment{}).Error; err != nil {
		log.Fatalln("删除评论时失败", err)
		return
	}

}
