package models

import "time"

type MODEL struct {
	ID				uint			`gorm:"primarykey" json:"id"`	// 主键id
	CreatedAt	time.Time	`json:"created_at"`						// 创建时间
	UpdatedAt time.Time	`json:"-"`										// 更新时间
}
// 删除图片
type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}
// 页数
type PageInfo struct {
	Page  int		 `form:"page"`
	Key		string `form:"key"`
	Limit int		 `form:"limit"`
	Sort  string `form:"sort"`
}
