package global

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type R struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var (
	Db     *gorm.DB
	Logger *logrus.Logger
)

func NewError(code int, msg string) *R {
	return &R{
		Code: code,
		Msg:  msg,
	}
}

// 实现error接口
func (e *R) Error() string {
	return e.Msg
}
