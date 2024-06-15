package service

import (
	"foodV5/common/dto"
	"foodV5/common/entity"
	"foodV5/common/repo"
	"gorm.io/gorm"
)

type OrderService struct {
	Db        *gorm.DB
	OrderRepo *repo.OrderRepo
}

// List 获取用户订单列表
func (os *OrderService) List(userId int64, pagedDto dto.Dto) (order []*entity.Order, err error) {
	return os.OrderRepo.List(userId, pagedDto)
}
