package middleware

import (
	"fmt"
	"foodV5/common/pkg/logs"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (m *Middleware) BaseMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		for key, values := range ctx.Request.Header {
			fmt.Println(key+"->", values)
		}
		logs.Log.Infof("用户ip:[%s],访问模式:[%s],接口路径:[%v]", ctx.Request.Header.Get("X-Forwarded-For"), ctx.Request.Method, ctx.Request.URL.Path)
		page, _ := strconv.Atoi(ctx.DefaultPostForm("page", "1"))
		limit, _ := strconv.Atoi(ctx.DefaultPostForm("limit", "30"))
		ctx.Set("page", page)
		ctx.Set("limit", limit)
		ctx.Next()
	}
}
