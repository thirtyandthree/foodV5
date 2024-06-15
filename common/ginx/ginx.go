package ginx

import (
	"foodV5/common/pkg/jwt"
	"foodV5/common/pkg/msg"
	"github.com/gin-gonic/gin"
)

func GetUserId(c *gin.Context) (userId int64) {
	token := c.GetHeader("Authorization")
	if token == "" {
		return
	}
	var err error
	if userId, err = jwt.GetUserIdByToken(token); err != nil {
		return
	}
	return
}

func ResponseData(c *gin.Context, data interface{}) {
	msg.Success(c, "success", data)
}

func ResponseFail(c *gin.Context, err interface{}) {
	msg.Fail(c, err, nil)
}

func ResponseDataFail(c *gin.Context, message interface{}, err error) {
	msg.Fail(c, message, err)
}

func ResponseSuccess(c *gin.Context, message string, data interface{}) {
	msg.Success(c, message, data)
}

func ResponseJson(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, data)
}

func ResponseXml(ctx *gin.Context, data interface{}) {
	ctx.XML(200, data)
}
