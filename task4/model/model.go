package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null" json:"-"`
	Email    string `gorm:"unique;not null" json:"email"`
}

type Post struct {
	gorm.Model
	Title   string `gorm:"not null" json:"title" binding:"required" msg:"文章标题不能为空"`
	Content string `gorm:"not null" json:"content" binding:"required" msg:"文章内容不能为空"`
	UserID  uint
	User    User
}

type Comment struct {
	gorm.Model
	Content string `gorm:"not null" json:"content" binding:"required" msg:"评论内容不能为空"`
	//这里的userid是评论人不是文章作者
	UserID uint
	User   User
	PostID uint `gorm:"not null" json:"postId" binding:"required" msg:"评论文章id不能为空"`
	Post   Post
}
