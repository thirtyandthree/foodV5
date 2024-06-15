package entity

type Time struct {

	// 创建时间
	CreateTime string `json:"create_time" gorm:"size:200"`
	UpdateTime string `json:"update_time" gorm:"size:200"`
	// 修改时间
	// 年月日
	Year  uint16 `json:"year"`
	Month uint16 `json:"month"`
	Day   uint16 `json:"day"`
}
