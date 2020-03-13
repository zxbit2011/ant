package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/zxbit2011/ant/manage-api/global"
	"github.com/zxbit2011/ant/manage-api/handle"
	"github.com/zxbit2011/ant/manage-api/handle/file"
	"github.com/zxbit2011/ant/manage-api/handle/sso"
	"github.com/zxbit2011/ant/manage-api/router"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			println("异常：%v", err)
			time.Sleep(1 * time.Minute)
		}
	}()
	err := global.InitGlobal(&global.ConfPath{
		ConfigPath: "config/config.toml",
	})
	if err != nil {
		println("初始化失败：", err.Error())
		time.Sleep(1 * time.Minute)
		return
	}

	global.Log.Info("开始启动")

	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	//Gzip压缩
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: DefaultSkipper,
		Level:   5,
	}))
	e.Use(middleware.Secure())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("6a2t6WKgSl75N40SFOhglShSJ0ua8OnZ"))))
	//文件清单
	e.Static("/files", "files")
	e.GET("/files/*/auth/*", func(c echo.Context) error {
		return c.File(strings.TrimLeft(c.Request().URL.Path, "/"))
	}, handle.Filter)
	group := e.Group("/api")
	{
		//登录2
		group.GET("/img/code", sso.ImageCode)
		//账号登录
		group.POST("/login", sso.Login)
		auth := group.Group("/auth", handle.Filter)
		{
			auth.POST("/logout", sso.LogOut)
			auth.POST("/login/menu", sso.GetLoginMenu)
			auth.POST("/files/upload/:auth", file.UploadFile)
			auth.POST("/files/remove", file.DelFileLog)
			router.SysNotifyRouter(auth)
			router.SysRouter(auth)
		}
	}
	//for _,v:=range e.Routes(){
	//	println(v.Path)
	//}
	go func() {
		if err := e.Start(fmt.Sprintf(":%d", global.Conf.HTTP.Port)); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}

//排除动态资源，如：图形验证码
func DefaultSkipper(c echo.Context) bool {
	if c.Path() == "/api/img/code" {
		return true
	}
	return false
}
