package dao

import (
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

const tmUserTableName = "tm_user"

// GetTmUserByPages TmUserInterfaceImp.GetTmUser 查询tm_user表里面的所有用户信息
func (t TmUserInterfaceImp) GetTmUserByPages(page int, pageSize int) (model.TmUserPage, error) {
	var err error
	var tmUserPage = model.TmUserPage{Page: page, PageSize: pageSize}

	cli := db.Get()
	offset := (page - 1) * pageSize
	err = cli.Table(tmUserTableName).
		Order("id asc").
		Limit(pageSize).Offset(offset).Scan(&tmUserPage.Data).
		Limit(-1).Offset(-1).
		Count(&tmUserPage.Total).Error

	return tmUserPage, err

}

func (t TmUserInterfaceImp) GetAllTmUser() ([]model.TmUserModel, error) {
	var err error
	var tmUsers = make([]model.TmUserModel, 0)
	cli := db.Get()
	err = cli.Table(tmUserTableName).Find(&tmUsers).Error
	return tmUsers, err
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

	err = t.DeleteTmUser(user.OpenId)
	if err != nil {
		return err
	}

	user.Id = 0
	err = UserImp.SaveUser(user)
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

func (t TmUserInterfaceImp) DeleteTmUser(openid string) error {
	//TODO implement me
	var err error
	cli := db.Get()
	err = cli.Table(tmUserTableName).Where("openid=?", openid).Delete(&model.TmUserModel{}).Error
	return err
}
