package msg

import (
	"foodV5/common/pkg/errors"
	"github.com/gin-gonic/gin"
)

// Msg 处理的结构体
type Msg struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (e *Msg) Error() string {
	return e.Msg
}

func Message(Code int, msg string, data interface{}) *Msg {
	if data == nil {
		data = ""
	}
	return &Msg{
		Code: Code,
		Msg:  msg,
		Data: data,
	}
}

func Success(ctx *gin.Context, message string, data interface{}) {
	ctx.JSON(200, Message(errors.Success, message, data))
}

func Fail(ctx *gin.Context, message interface{}, data interface{}) {
	code := errors.BusinessCode
	msg := ""
	switch m := message.(type) {
	case string:
		msg = m
	case *errors.Err:
		code = message.(*errors.Err).Code
		msg = message.(*errors.Err).Msg
	case error:
		msg = message.(error).Error()
	}
	ctx.JSON(200, Message(code, msg, data))
}
