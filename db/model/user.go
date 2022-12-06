package model

// UserModel 实际用户模型
type UserModel struct {
	Id         int32  `gorm:"column:id" json:"id"`
	OpenId     string `gorm:"column:openid" json:"openid"`
	Permission string `gorm:"column:permission" json:"permission"`
	IsAdmin    bool   `gorm:"column:isAdmin" json:"isAdmin"`
}

// Pages 分页存储用户模型
type Pages struct {
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
	Total    int64       `json:"total"`
	Data     []UserModel `json:"data"`
}
