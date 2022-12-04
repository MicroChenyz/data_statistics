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
	//TODO implement me
	panic("implement me")
}

func (u UserInterfaceImp) ClearUser(id int32) error {
	//TODO implement me
	panic("implement me")
}
