package repo

import (
	"foodV5/common/entity"
	"gorm.io/gorm"
)

type InviteRepo struct {
	DB   *gorm.DB
	Repo *Repo
}

func (i *InviteRepo) CreateWithTx(db *gorm.DB, invite *entity.Invite) error {
	i.Repo.CreateTime(&invite.Time)
	return db.Create(invite).Error
}

func (i *InviteRepo) FindByTo(toUser int64) (invite *entity.Invite, err error) {
	err = i.DB.Where("to_user=?", toUser).Find(&invite).Error
	return
}

// FindInviteCount 查询邀请好友量
func (i *InviteRepo) FindInviteCount(userId int64, level string) (c int64, err error) {
	// 一级邀请还是
	if level == "one" {
		err = i.DB.Model(&entity.Invite{}).
			Where("from_user=?", userId).
			Count(&c).Error
	} else {
		err = i.DB.Model(&entity.Invite{}).
			Where("level_one=?", userId).
			Count(&c).Error
	}
	return
}
