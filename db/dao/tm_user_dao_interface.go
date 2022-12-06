package dao

import "wxcloudrun-golang/db/model"

// TmUserInterface 暂存用户模型接口
type TmUserInterface interface {
	GetTmUserByPages(page int, pageSize int) (model.TmUserPage, error)
	GetAllTmUser() ([]model.TmUserModel, error)
	SaveTmUser(tmUser *model.TmUserModel) error
	UpdateTmUser(user *model.UserModel) error
	GetTmUserByOpenid(openid string) (model.TmUserModel, error)
	DeleteTmUser(openid string) error
}

type TmUserInterfaceImp struct{}

var TmUserImp TmUserInterface = &TmUserInterfaceImp{}
