package dao

import (
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

const userTableName = "user"

func (u UserInterfaceImp) SaveUser(user *model.UserModel) error {
	//TODO implement me
	var err error
	cli := db.Get()
	err = cli.Table(userTableName).Save(user).Error
	return err
}

func (u UserInterfaceImp) FindUserByOpenId(openId string) (model.UserModel, error) {
	var err error
	var user model.UserModel
	cli := db.Get()
	err = cli.Table(userTableName).Where("openid=?", openId).First(&user).Error
	return user, err
}

func (u UserInterfaceImp) ClearUser(id int) error {
	//TODO implement me
	var err error
	cli := db.Get()
	err = cli.Table(userTableName).Delete(&model.UserModel{Id: id}).Error
	return err
}

func (u UserInterfaceImp) FindAllUser(page int, pageSize int) (model.Pages, error) {
	//TODO implement me
	var err error
	var pages = model.Pages{Page: page, PageSize: pageSize}

	cli := db.Get()
	offset := (page - 1) * pageSize
	err = cli.Table(userTableName).
		Order("id asc").
		Limit(pageSize).
		Offset(offset).
		Scan(&pages.Data).
		Limit(-1).Offset(-1).Count(&pages.Total).Error

	return pages, err
}
