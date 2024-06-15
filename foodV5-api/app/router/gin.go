package router

import (
	"github.com/gin-gonic/gin"
)

func NewGin() *gin.Engine {

	app := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	return app
}
