package middleware

import (
	"foodV5/common/ginx"
	"github.com/gin-gonic/gin"
)

func (m *Middleware) RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ginx.ResponseFail(ctx, err)
				ctx.IsAborted()
			}
		}()
		ctx.Next()
	}
}
