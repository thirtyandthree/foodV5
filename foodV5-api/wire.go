//go:build wireinject
// +build wireinject

package main

import (
	"foodV5/common/middleware"
	"foodV5/common/pkg/wechat"
	"foodV5/common/repo"
	"foodV5/common/service"
	"foodV5/foodV5-api/app"
	"foodV5/foodV5-api/app/api"
	"foodV5/foodV5-api/app/router"
	service2 "foodV5/foodV5-api/app/service"
	"github.com/google/wire"
)

func wireApp() (*app.Application, error) {
	panic(
		wire.Build(
			app.NewGorm,
			app.NewRedis,
			router.NewGin,
			wechat.NewWechatMini,
			router.WireRouter,
			app.WireApp,
			api.WireController,
			middleware.WireMiddleware,
			repo.WireCommonRepo,
			service.WireCommonService,
			service2.WireService,
		))
}
