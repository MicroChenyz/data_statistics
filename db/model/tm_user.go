package model

import "time"

// TmUserModel 暂存用户模型
type TmUserModel struct {
	Id       int32     `gorm:"column:id" json:"id"`
	OpenId   string    `gorm:"column:openid" json:"openid"`
	Username string    `gorm:"column:username" json:"username"`
	CreateAt time.Time `gorm:"column:createAt" json:"createAt"`
}
