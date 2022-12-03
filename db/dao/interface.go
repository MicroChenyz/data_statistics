package dao

import "wxcloudrun-golang/db/model"

// TmUserInterface 暂存用户模型接口
type TmUserInterface interface {
	GetTmUser() ([]model.TmUserModel, error)
	SaveTmUser(tmUser *model.TmUserModel) error
	ClearTmUser(id int32) error
}

type TmUserInterfaceImp struct{}

var TmUserImp TmUserInterface = &TmUserInterfaceImp{}
