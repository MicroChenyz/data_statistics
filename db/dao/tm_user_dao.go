package dao

import (
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

const tmUserTableName = "tm_user"

// GetTmUser TmUserInterfaceImp.GetTmUser 查询tm_user表里面的所有用户信息
func (t TmUserInterfaceImp) GetTmUser() ([]model.TmUserModel, error) {
	var err error
	var tmUserSlice = make([]model.TmUserModel, 0)

	cli := db.Get()
	err = cli.Table(tmUserTableName).Find(&tmUserSlice).Error

	return tmUserSlice, err

}

// SaveTmUser TmUserInterfaceImp.SaveTmUser 存储一条用户信息
func (t TmUserInterfaceImp) SaveTmUser(tmUser *model.TmUserModel) error {
	var err error
	cli := db.Get()
	err = cli.Table(tmUserTableName).Save(tmUser).Error
	return err

}

func (t TmUserInterfaceImp) ClearTmUser(user *model.UserModel) error {
	var err error
	cli := db.Get()
	tx := cli.Begin()
	defer tx.Rollback()

	if err := tx.Error; err != nil {
		return err
	}

	_, err = t.GetTmUserByOpenid(user.OpenId)
	if err != nil {
		return err
	}

	err = tx.Table(tmUserTableName).Delete(&model.TmUserModel{Id: user.Id}).Error
	if err != nil {
		return err
	}

	user.Id = 0
	err = tx.Table(userTableName).Save(user).Error
	if err != nil {
		return err
	}

	return tx.Commit().Error
}

func (t TmUserInterfaceImp) GetTmUserByOpenid(openid string) (model.TmUserModel, error) {
	var err error
	var tmUser model.TmUserModel
	cli := db.Get()
	err = cli.Table(tmUserTableName).Where("openid=?", openid).First(&tmUser).Error
	return tmUser, err
}
