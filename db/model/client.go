package model

import "time"

// Client 客户数据模型
type Client struct {
	Id           int32     `gorm:"column:id" json:"id"`
	ClientName   string    `gorm:"column:client_name" json:"client_name"`
	StoveNum     string    `gorm:"column:stove_num" json:"stove_num"`
	StoveCap     float64   `gorm:"column:stove_cap" json:"stove_cap"`
	IronNotch    string    `gorm:"column:iron_notch" json:"iron_notch"`
	IronNotchNum int       `gorm:"column:iron_notch_num" json:"iron_notch_num"`
	MudType      string    `gorm:"column:mud_type" json:"mud_type"`
	CreateTime   time.Time `gorm:"column:create_time" json:"create_time"`
}

type ClientResponse struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}
