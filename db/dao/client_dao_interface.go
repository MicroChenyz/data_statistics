package dao

import "wxcloudrun-golang/db/model"

// ClientInterface 客户模型接口
type ClientInterface interface {
	SaveClient(client *model.Client) error
	DeleteClientByClientId(id int) error
	DeleteClientByOpenId(openid string) error
	GetClientByOpenId(openid string) ([]model.ClientResponse, error)
	GetClientByClientId(id int) (model.Client, error)
	GetAllClientByPages(page int, pageSize int) (model.ClientPage, error)
	GetAllClient() ([]model.Client, error)
}

type ClientInterfaceImp struct{}

var ClientImp ClientInterface = &ClientInterfaceImp{}
