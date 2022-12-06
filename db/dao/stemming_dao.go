package dao

import (
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

const stemmingTableName = "stemming"

func (s StemmingInterfaceImp) SaveStemming(stemming *model.Stemming) error {
	//TODO implement me
	var err error
	cli := db.Get()
	err = cli.Table(stemmingTableName).Save(stemming).Error
	return err
}

func (s StemmingInterfaceImp) GetStemmingByOpenid(openid string) ([]model.StemmingResponse, error) {
	//TODO implement me
	var err error
	var stemmingResponse = make([]model.StemmingResponse, 0)
	var stemmings = make([]model.Stemming, 0)
	cli := db.Get()
	err = cli.Table(stemmingTableName).Where("openid=?", openid).Find(&stemmings).Error
	if err != nil {
		return nil, err
	}
	for i := range stemmings {
		var stemmingRes model.StemmingResponse
		stemmingRes.Id = stemmings[i].Id
		stemmingRes.StartTime = stemmings[i].StartTime
		stemmingRes.EndTime = stemmings[i].EndTime
		stemmingRes.CreateTime = stemmings[i].CreateTime
		stemmingRes.ModifyTime = stemmings[i].ModifyTime
		stemmingResponse = append(stemmingResponse, stemmingRes)
	}
	return stemmingResponse, err
}

func (s StemmingInterfaceImp) GetStemmingById(id int) (model.Stemming, error) {
	//TODO implement me
	var err error
	var stemming = model.Stemming{}
	cli := db.Get()
	err = cli.Table(stemmingTableName).First(&stemming, id).Error
	return stemming, err
}

func (s StemmingInterfaceImp) UpdateStemmingById(stemming *model.Stemming) error {
	//TODO implement me
	var err error
	cli := db.Get()
	err = cli.Table(stemmingTableName).Updates(stemming).Error
	return err
}
