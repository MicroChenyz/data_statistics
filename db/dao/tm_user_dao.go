package dao

import (
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

const tmUserTableName = "tm_user"

// GetTmUserByPages TmUserInterfaceImp.GetTmUser 查询tm_user表里面的所有用户信息
func (t TmUserInterfaceImp) GetTmUserByPages(page int, pageSize int) (model.TmUserPage, error) {
	var err error
	var tmUserPage model.TmUserPage

	cli := db.Get()
	offset := (page - 1) * pageSize
	err = cli.Table(tmUserTableName).
		Order("id asc").
		Limit(pageSize).Offset(offset).Scan(&tmUserPage.Data).
		Limit(-1).Offset(-1).
		Count(&tmUserPage.Total).Error

	return tmUserPage, err

}

// SaveTmUser TmUserInterfaceImp.SaveTmUser 存储一条用户信息
func (t TmUserInterfaceImp) SaveTmUser(tmUser *model.TmUserModel) error {
	var err error
	cli := db.Get()
	err = cli.Table(tmUserTableName).Save(tmUser).Error
	return err

}

func (t TmUserInterfaceImp) UpdateTmUser(user *model.UserModel) error {
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
