package model

// UserModel 实际用户模型
type UserModel struct {
	Id         int32  `gorm:"column:id" json:"id"`
	OpenId     string `gorm:"column:openid" json:"openid"`
	Permission string `gorm:"column:permission" json:"permission"`
}
