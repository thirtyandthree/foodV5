package app

import (
	"context"
	"fmt"
	"foodV5/common/config"
	"foodV5/foodV5-api/app/router"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Application struct {
	Router *router.Router
	Http   *http.Server `wire:"-"`
}

func (app *Application) Init() error {
	var cstZone = time.FixedZone("CST", 8*3600)
	time.Local = cstZone

	return nil
}

func (app *Application) HttpServerStart() {
	cfg := config.C.Server
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	// 初始化中间件
	app.Router.InitMiddleware()
	// 不用授权登陆的
	app.Router.NoAuth()
	// 需要授权登陆的
	app.Router.Auth()

	app.Http = &http.Server{
		Addr:    addr,
		Handler: app.Router.Engine,
	}
	go func() {
		if err := app.Http.ListenAndServe(); err != nil {
			fmt.Printf("Http服务启动异常: %v\n", err)
			return
		}
	}()
}

func (app *Application) Run() {
	if err := app.Init(); err != nil {
		return
	}

	app.HttpServerStart()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	app.Stop()
}

func (app *Application) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.Http.Shutdown(ctx); err != nil {
		fmt.Printf("Http 服务关闭错误:%v\n", err)
		return
	}
	fmt.Println("Http server 关闭成功")
}
