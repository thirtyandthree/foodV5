package middleware

import (
	"foodV5/common/config"
	"foodV5/common/ginx"
	"github.com/gin-gonic/gin"
)

func (m *Middleware) CloseMiddleware() gin.HandlerFunc {
	// 检查是否开放站点
	return func(context *gin.Context) {
		if config.C.Server.IsClose {
			ginx.ResponseFail(context, config.C.Server.CloseMessage)
			context.Abort()
		}
		context.Next()
	}

}
