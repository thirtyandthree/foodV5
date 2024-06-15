package entity

// Invite 邀请表
type Invite struct {
	Id int64 `json:"id" gorm:"primaryKey,autoIncrement"`
	// 邀请人
	FromUser int64 `json:"from_user"`
	// 被邀请人
	ToUser int64 `json:"to_user"`
	// 是否奖励
	IsReward uint8 `json:"is_reward"`
	// 是否已奖励
	Amount float64 `json:"amount"`
	// 等级
	LevelOne int64 `json:"level_one" gorm:"type:int(11);comment:顶级邀请人:三级分销里面第一位邀请用户"`

	Time
}
