package dao

import "wxcloudrun-golang/db/model"

// UserInterface 实际用户模型接口
type UserInterface interface {
	SaveUser(user *model.UserModel) error
	FindUserByOpenId(openId string) (model.UserModel, error)
	FindUserByPages(page int, pageSize int) (model.UserPage, error)
	FindAllUser() ([]model.UserModel, error)
	DeleteUser(openid string) error
}

type UserInterfaceImp struct{}

var UserImp UserInterface = &UserInterfaceImp{}
