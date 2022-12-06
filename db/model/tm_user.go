package model

import "time"

// TmUserModel 暂存用户模型
type TmUserModel struct {
	Id       int       `gorm:"column:id" json:"id"`
	OpenId   string    `gorm:"column:openid" json:"openid"`
	Username string    `gorm:"column:username" json:"username"`
	CreateAt time.Time `gorm:"column:createAt" json:"createAt"`
}

type TmUserPage struct {
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
	Total    int64       `json:"total"`
	Data     []UserModel `json:"data"`
}
