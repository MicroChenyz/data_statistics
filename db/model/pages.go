package model

type Pages struct {
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
	Total    int64       `json:"total"`
	Data     []UserModel `json:"data"`
}
