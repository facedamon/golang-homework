package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/facedamon/golang-homework/blog/config"
	"github.com/facedamon/golang-homework/blog/global"
	"github.com/facedamon/golang-homework/blog/model"
	"github.com/facedamon/golang-homework/blog/router"
)

func main() {
	config.InitConfig()
	config.InitLogrus()
	config.InitDB()
	if err := global.Db.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{}); err != nil {
		global.Logger.Debug("AutoMigrate失败")
		return
	}
	g := router.InitRouter()

	//err := g.Run(config.AppConfig.App.Port)
	//if err != nil {
	//	return
	//}

	srv := &http.Server{
		Addr:    config.AppConfig.App.Port,
		Handler: g,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Logger.Errorln("listen:", err)
		}
	}()
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no params) by default sends syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	global.Logger.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		global.Logger.Println("Server Shutdown:", err)
	}
	global.Logger.Println("Server exiting")
}
