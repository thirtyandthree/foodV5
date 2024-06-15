package middleware

import (
	"foodV5/common/ginx"
	"foodV5/common/pkg/errors"
	"github.com/gin-gonic/gin"
	"time"
)

func (m *Middleware) AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 检查redis是否有这个token,没有代表token无效
		token := ctx.GetHeader("Authorization")
		s, err := m.Client.Exists(ctx, token).Result()
		if err != nil {
			ginx.ResponseFail(ctx, err)
			ctx.Abort()
		}
		if s != 1 {
			ginx.ResponseFail(ctx, errors.TokenError)
			ctx.Abort()
			return
		}
		// 从上下文根据传递来的token获取对应的用户id
		userId := ginx.GetUserId(ctx)
		// id小于0，就是token有误相当于,就驳回，不准请求
		if userId <= 0 {
			err := errors.UserUnLoginError
			ginx.ResponseFail(ctx, err)
			ctx.Abort()
		}
		// 用户耍小聪明，瞎几把乱搞,实际上我没有这个用户呢?

		// 设置uid为用户的id
		ctx.Set("uid", userId)
		// 设置超时时间,三天用请求就过期
		m.Client.Expire(ctx, token, time.Second*60*60*24*3)
		// 放行请求
		ctx.Next()
	}
}
