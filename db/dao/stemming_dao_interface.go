package dao

import "wxcloudrun-golang/db/model"

type StemmingInterface interface {
	SaveStemming(stemming *model.Stemming) error
	GetStemmingByOpenid(openid string) ([]model.StemmingResponse, error)
	GetStemmingById(id int) (model.Stemming, error)
	UpdateStemmingById(stemming *model.Stemming) error
}

type StemmingInterfaceImp struct{}

var StemmingImp StemmingInterface = &StemmingInterfaceImp{}
