package service

import (
	"foodV5/common/entity"
	"foodV5/common/repo"
	"gorm.io/gorm"
)

// InviteService 邀请服务
type InviteService struct {
	InviteRepo *repo.InviteRepo
}

// Plus 新增邀请记录
func (i *InviteService) Plus(tx *gorm.DB, invite *entity.Invite) error {
	return i.InviteRepo.CreateWithTx(tx, invite)
}
