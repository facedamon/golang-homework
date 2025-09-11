package config

import (
	"log"
	"time"

	"github.com/facedamon/golang-homework/gorm_test/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {
	db, err := gorm.Open(mysql.Open(AppConfig.Database.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("failed to init db", err)
		return
	}
	sqlDb, err := db.DB()
	if err != nil {
		log.Fatalln("failed to get db", err)
		return
	}
	sqlDb.SetMaxIdleConns(AppConfig.Database.MaxIdleConns)
	sqlDb.SetMaxOpenConns(AppConfig.Database.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Hour)

	global.Db = db
}
