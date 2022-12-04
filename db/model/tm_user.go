package model

import "time"

// TmUserModel 暂存用户模型
type TmUserModel struct {
	Id       int32     `gorm:"column:id" json:"id"`
	OpenId   string    `gorm:"column:openid" json:"openid"`
	TelPhone string    `gorm:"column:telPhone" json:"telPhone"`
	CreateAt time.Time `gorm:"column:createAt" json:"createAt"`
	Remark   string    `gorm:"column:remark" json:"remark"`
}
