package dao

import "wxcloudrun-golang/db/model"

// UserInterface 实际用户模型接口
type UserInterface interface {
	SaveUser(user *model.UserModel) error
	FindUserByOpenId(openId string) (model.UserModel, error)
	FindAllUser(page int, pageSize int) (model.Pages, error)
	ClearUser(id int32) error
}

type UserInterfaceImp struct{}

var UserImp UserInterface = &UserInterfaceImp{}
