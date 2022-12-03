package dao

import (
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

const tableName = "tm_user"

// GetTmUser TmUserInterfaceImp.GetTmUser 查询tm_user表里面的所有用户信息
func (t TmUserInterfaceImp) GetTmUser() ([]model.TmUserModel, error) {
	//TODO implement me
	var err error
	var tmUserSlice = make([]model.TmUserModel, 0)

	cli := db.Get()
	err = cli.Table(tableName).Find(&tmUserSlice).Error

	return tmUserSlice, err

}

// SaveTmUser TmUserInterfaceImp.SaveTmUser 存储一条用户信息
func (t TmUserInterfaceImp) SaveTmUser(tmUser *model.TmUserModel) error {
	var err error
	cli := db.Get()
	err = cli.Table(tableName).Save(tmUser).Error
	return err

}

func (t TmUserInterfaceImp) ClearTmUser(id int32) error {
	var err error
	cli := db.Get()
	err = cli.Table(tableName).Delete(&model.TmUserModel{Id: id}).Error
	return err
}
