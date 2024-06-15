package api

import (
	"fmt"
	"foodV5/common/dto"
	"foodV5/common/ginx"
	"github.com/gin-gonic/gin"
)

// OrderApi 订单API接口
type OrderApi struct {
}

// List 获取用户订单列表
func (o *OrderApi) List(ctx *gin.Context) {
	// 1.0 获取用户的id
	userId := ctx.GetInt64("uid")
	pageDto := &dto.Dto{}
	if err := ctx.ShouldBindJSON(pageDto); err != nil {
		ginx.ResponseDataFail(ctx, "获取订单列表失败", nil)
		return
	}
	// 查询列表分页
	fmt.Println(userId)

}
