package middleware

import (
	"foodV5/common/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

//func (m *Middleware) CORSMiddleware() gin.HandlerFunc {
//	return func(context *gin.Context) {
//		method := context.Request.Method
//		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
//		context.Header("Access-Control-Allow-Origin", "*") // 设置允许访问所有域
//		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
//		context.Header("Access-Control-Allow-Headers", "Authorization,Authorizations, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma,token,openid,opentoken")
//		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
//		context.Header("Access-Control-Max-Age", "172800")
//		context.Header("Access-Control-Allow-Credentials", "false")
//		context.Set("content-type", "application/json")
//		if method == "OPTIONS" {
//			context.JSON(http.StatusOK, nil)
//		}
//
//		//处理请求
//		context.Next()
//	}
//}

func (m *Middleware) CORSMiddleware() gin.HandlerFunc {
	cfg := config.C.CORS
	return cors.New(cors.Config{
		AllowOrigins:     cfg.AllowOrigins,
		AllowMethods:     cfg.AllowMethods,
		AllowHeaders:     cfg.AllowHeaders,
		AllowCredentials: cfg.AllowCredentials,
		MaxAge:           time.Second * time.Duration(cfg.MaxAge),
	})
}
