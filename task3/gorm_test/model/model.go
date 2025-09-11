package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `gorm:"name"`
	Posts []Post
	Num   *int `gorm:"num"`
}

type Post struct {
	gorm.Model
	Title     string `gorm:"title"`
	Content   string `gorm:"content"`
	UserID    uint
	Comments  []Comment
	ComStatus string `gorm:"com_status"`
}

type Comment struct {
	gorm.Model
	PostID uint   `gorm:"post_id"`
	Comm   string `gorm:"comm"`
}

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	if err := tx.Debug().Model(&User{}).Where("id=?", p.UserID).UpdateColumn("num", gorm.Expr("num + ?", 1)).Error; err != nil {
		return err
	}
	return nil
}

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var count int64
	if err := tx.Debug().Model(&Post{}).Where("post_id=?", c.PostID).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		if err = tx.Debug().Model(&Post{}).UpdateColumn("com_status", "无评论").Where("id=?", c.PostID).Error; err != nil {
			return err
		}
	}
	return nil
}
