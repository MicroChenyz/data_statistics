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

func (u UserInterfaceImp) DeleteUser(openid string) error {
	//TODO implement me
	var err error
	cli := db.Get()
	err = cli.Table(userTableName).Where("openid=?", openid).Delete(&model.UserModel{}).Error
	return err
}

func (u UserInterfaceImp) FindUserByPages(page int, pageSize int) (model.UserPage, error) {
	//TODO implement me
	var err error
	var pages = model.UserPage{Page: page, PageSize: pageSize}

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

func (u UserInterfaceImp) FindAllUser() ([]model.UserModel, error) {
	var err error
	var user = make([]model.UserModel, 0)

	cli := db.Get()
	err = cli.Table(userTableName).Find(&user).Error

	return user, err
}
