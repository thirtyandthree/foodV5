package api

import (
	"foodV5/common/ginx"
	"foodV5/common/pkg/invite_code"
	"foodV5/foodV5-api/app/dto"
	"foodV5/foodV5-api/app/service"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type AccountApi struct {
	AccountService *service.AccountService
}

// Login 小程序端用户进行登陆端接口
func (a *AccountApi) Login(ctx *gin.Context) {
	// 获取用户的code这些以及邀请码
	account := &dto.AccountDto{}
	if err := ctx.ShouldBindJSON(account); err != nil {
		ginx.ResponseFail(ctx, err)
		return
	}
	// 邀请码
	inviteFrom, _ := strconv.ParseInt(account.From, 10, 64)
	if inviteFrom <= 0 {
		inviteFrom = invite_code.Decode(account.From)
	}
	// 进行小程序登陆
	token, err := a.AccountService.MiniLogin(account.Code, inviteFrom)
	if err != nil {
		ginx.ResponseFail(ctx, err.Error())
		return
	}
	a.AccountService.SetTokenInfo(token, time.Second*60*60*24*3)
	ginx.ResponseSuccess(ctx, "登陆成功", map[string]interface{}{
		"token": token,
	})
}
