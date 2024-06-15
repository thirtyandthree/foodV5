package router

import (
	"foodV5/common/config"
	"foodV5/common/ginx"
	"foodV5/common/middleware"
	"foodV5/foodV5-api/app/api"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine     *gin.Engine
	Middleware *middleware.Middleware

	AccountApi *api.AccountApi
}

// InitMiddleware 加载中间件
func (r *Router) InitMiddleware() {
	r.Engine.Use(r.Middleware.BaseMiddleware())
	r.Engine.Use(r.Middleware.RecoveryMiddleware())
	// 站点是否开放
	if config.C.Server.IsClose {
		r.Engine.Use(r.Middleware.CloseMiddleware())
	}
	// 跨域配置
	if config.C.CORS.Enable {
		r.Engine.Use(r.Middleware.CORSMiddleware())
	}
}

// NoAuth 不需要授权的页面
func (r *Router) NoAuth() {
	r.Engine.NoRoute(func(ctx *gin.Context) {
		ginx.ResponseDataFail(ctx, "你访问尼玛呢", nil)
	})
	account := r.Engine.Group("account")
	{
		account.POST("login", r.AccountApi.Login)
	}
}

// Auth 需要授权的页面
func (r *Router) Auth() {
}
