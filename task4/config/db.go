package config

import (
	"time"

	"github.com/facedamon/golang-homework/blog/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() {
	myWriter := &gormWriter{global.Logger}
	gormLogger := logger.New(myWriter, logger.Config{
		SlowThreshold: time.Millisecond * 100,
		LogLevel:      logger.Info,
	})

	db, err := gorm.Open(mysql.Open(AppConfig.Database.Dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		global.Logger.Errorln("failed to init db", err)
		return
	}

	sqlDb, err := db.DB()
	if err != nil {
		global.Logger.Errorln("failed to get db", err)
		return
	}
	sqlDb.SetMaxIdleConns(AppConfig.Database.MaxIdleConns)
	sqlDb.SetMaxOpenConns(AppConfig.Database.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Hour)

	if err = sqlDb.Ping(); err != nil {
		global.Logger.Errorln("mysql ping failed.", err)
		return
	}
	global.Logger.Println("mysql ping successfully.")
	global.Db = db
}
