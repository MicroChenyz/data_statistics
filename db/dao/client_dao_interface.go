package dao

import "wxcloudrun-golang/db/model"

// ClientInterface 客户模型接口
type ClientInterface interface {
	SaveClient(client *model.Client) error
	ClearClient(id int32) error
	GetClientByUserid(userId int32) ([]model.ClientResponse, error)
	GetClientById(id int32) (model.Client, error)
}

type ClientInterfaceImp struct{}

var ClientImp ClientInterface = &ClientInterfaceImp{}
