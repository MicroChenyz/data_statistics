package dao

import "wxcloudrun-golang/db/model"

// TmUserInterface 暂存用户模型接口
type TmUserInterface interface {
	GetTmUser() ([]model.TmUserModel, error)
	SaveTmUser(tmUser *model.TmUserModel) error
	ClearTmUser(user *model.UserModel) error
	GetTmUserByOpenid(openid string) (model.TmUserModel, error)
}

type TmUserInterfaceImp struct{}

var TmUserImp TmUserInterface = &TmUserInterfaceImp{}
