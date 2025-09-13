package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Dsn          string
		MaxIdleConns int
		MaxOpenConns int
	}
	App struct {
		Port string
	}
}

var AppConfig Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("reading config file", err)
		return
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatalln("unmarshal config file", err)
		return
	}
}
